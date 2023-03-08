package main

import (
	"fmt"
	"net"
)

func isPortFree(port uint16) error {
	listen, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		return err
	}
	_ = listen.Close()
	return nil
}
