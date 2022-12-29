package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"runtime/debug"
	"sync"
	"syscall"
	"time"

	corev1 "k8s.io/api/core/v1"

	"github.com/alexec/kit/internal/proc"
	"github.com/alexec/kit/internal/types"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"sigs.k8s.io/yaml"
)

func up() *cobra.Command {
	cmd := &cobra.Command{
		Use: "up",
		RunE: func(cmd *cobra.Command, args []string) error {

			ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGTERM)
			defer stop()

			_ = os.Mkdir("logs", 0777)

			pod := &types.Pod{}
			status := &types.Status{}
			logEntries := make(map[string]*types.LogEntry)
			init := true

			go func() {
				defer handleCrash(stop)
				for {
					width, _, _ := terminal.GetSize(0)
					if width == 0 {
						width = 80
					}

					log.Printf("%s[2J", escape)
					log.Printf("%s[H", escape)
					containers := pod.Spec.Containers
					if init {
						containers = pod.Spec.InitContainers
					}
					for _, c := range containers {
						state := status.GetContainerStatus(c.Name)
						if state == nil {
							continue
						}
						r := "▓"
						reason := ""
						if state.State.Waiting != nil {
							reason = state.State.Waiting.Reason
						} else if state.State.Running != nil {
							if state.Ready {
								r = color.GreenString("▓")
								reason = "ready"
							} else {
								r = color.BlueString("▓")
								reason = "running"
							}
						} else if state.State.Terminated != nil {
							reason = state.State.Terminated.Reason
							if reason == "error" {
								r = color.RedString("▓")
							}
						} else {
							reason = "unknown"
						}
						line := fmt.Sprintf("%s %-10s [%-7s]  %s", r, state.Name, reason, logEntries[c.Name].String())
						if len(line) > width && width > 0 {
							line = line[0 : width-1]
						}
						log.Println(line)
					}
					time.Sleep(time.Second / 2)
				}
			}()

			in, err := os.ReadFile(kitFile)
			if err != nil {
				return err
			}
			if err = yaml.UnmarshalStrict(in, pod); err != nil {
				return err
			}
			data, err := yaml.Marshal(pod)
			if err != nil {
				return err
			}
			if err = os.WriteFile(kitFile, data, 0o0644); err != nil {
				return err
			}

			for _, containers := range [][]types.Container{pod.Spec.InitContainers, pod.Spec.Containers} {
				wg := &sync.WaitGroup{}

				if init {
					for _, c := range pod.Spec.InitContainers {
						logEntries[c.Name] = &types.LogEntry{}
						status.InitContainerStatuses = append(status.InitContainerStatuses, corev1.ContainerStatus{
							Name: c.Name,
						})
					}
				} else {
					for _, c := range pod.Spec.Containers {
						logEntries[c.Name] = &types.LogEntry{}
						status.ContainerStatuses = append(status.ContainerStatuses, corev1.ContainerStatus{
							Name: c.Name,
						})
					}
				}

				for _, c := range containers {

					name := c.Name
					state := status.GetContainerStatus(name)

					logEntry := logEntries[name]

					logFile, err := os.Create(filepath.Join("logs", name+".log"))
					if err != nil {
						return err
					}
					stdout := io.MultiWriter(logFile, logEntry.Stdout())
					stderr := io.MultiWriter(logFile, logEntry.Stderr())
					pd := proc.New(c)

					if err = pd.Init(ctx); err != nil {
						return err
					}

					wg.Add(1)
					go func() {
						defer handleCrash(stop)
						defer wg.Done()
						for {
							select {
							case <-ctx.Done():
								return
							default:
								err := func() error {
									state.State = corev1.ContainerState{
										Waiting: &corev1.ContainerStateWaiting{Reason: "stopping"},
									}
									if err := pd.Stop(ctx, pod.Spec.GetTerminationGracePeriod()); err != nil {
										return fmt.Errorf("failed to stop: %v", err)
									}
									state.State = corev1.ContainerState{
										Waiting: &corev1.ContainerStateWaiting{Reason: "building"},
									}
									if err := pd.Build(ctx, stdout, stderr); err != nil {
										return fmt.Errorf("failed to build: %v", err)
									}
									state.State = corev1.ContainerState{
										Running: &corev1.ContainerStateRunning{},
									}
									if err := pd.Run(ctx, stdout, stderr); err != nil {
										return fmt.Errorf("failed to run: %v", err)
									}
									return nil
								}()
								if err != nil && err != context.Canceled {
									state.State = corev1.ContainerState{
										Terminated: &corev1.ContainerStateTerminated{Reason: "error"},
									}
									logEntry = &types.LogEntry{Level: "error", Msg: err.Error()}
								} else {
									state.State = corev1.ContainerState{
										Terminated: &corev1.ContainerStateTerminated{Reason: "exited"},
									}
									if init {
										return
									}
								}
							}
						}
					}()

					go func() {
						<-ctx.Done()
						if err := pd.Stop(context.Background(), pod.Spec.GetTerminationGracePeriod()); err != nil {
							logEntry = &types.LogEntry{Level: "error", Msg: fmt.Sprintf("failed to stop: %v", err)}
						}
					}()

					if probe := c.LivenessProbe; probe != nil {
						liveFunc := func(live bool, err error) {
							if !live {
								if err := pd.Stop(ctx, pod.Spec.GetTerminationGracePeriod()); err != nil {
									logEntry = &types.LogEntry{Level: "error", Msg: err.Error()}
								}
							}
						}
						go probeLoop(ctx, stop, *probe, liveFunc)
					}
					if probe := c.ReadinessProbe; probe != nil {
						readyFunc := func(ready bool, err error) {
							state.Ready = ready
							if err != nil {
								logEntry = &types.LogEntry{Level: "error", Msg: err.Error()}
							}
						}
						go probeLoop(ctx, stop, *probe, readyFunc)
					}
				}

				wg.Wait()

				init = false
			}
			return nil
		},
	}
	return cmd
}
func handleCrash(stop func()) {
	if r := recover(); r != nil {
		log.Println(r)
		debug.PrintStack()
		stop()
	}
}