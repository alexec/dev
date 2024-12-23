# PodSpec Schema

```txt
https://github.com/kitproj/kit/internal/types/pod#/$defs/PodSpec
```

Task is a unit of work that should be run.

| Abstract            | Extensible | Status         | Identifiable | Custom Properties | Additional Properties | Access Restrictions | Defined In                                                            |
| :------------------ | :--------- | :------------- | :----------- | :---------------- | :-------------------- | :------------------ | :-------------------------------------------------------------------- |
| Can be instantiated | No         | Unknown status | No           | Forbidden         | Forbidden             | none                | [pod.schema.json\*](../../out/pod.schema.json "open original schema") |

## PodSpec Type

`object` ([PodSpec](pod-defs-podspec.md))

# PodSpec Properties

| Property                                                        | Type      | Required | Nullable       | Defined by                                                                                                                                                                                  |
| :-------------------------------------------------------------- | :-------- | :------- | :------------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| [logLevel](#loglevel)                                           | `string`  | Optional | cannot be null | [Untitled schema](pod-defs-podspec-properties-loglevel.md "https://github.com/kitproj/kit/internal/types/pod#/$defs/PodSpec/properties/logLevel")                                           |
| [terminationGracePeriodSeconds](#terminationgraceperiodseconds) | `integer` | Optional | cannot be null | [Untitled schema](pod-defs-podspec-properties-terminationgraceperiodseconds.md "https://github.com/kitproj/kit/internal/types/pod#/$defs/PodSpec/properties/terminationGracePeriodSeconds") |
| [tasks](#tasks)                                                 | `array`   | Optional | cannot be null | [Untitled schema](pod-defs-tasks.md "https://github.com/kitproj/kit/internal/types/pod#/$defs/PodSpec/properties/tasks")                                                                    |
| [volumes](#volumes)                                             | `array`   | Optional | cannot be null | [Untitled schema](pod-defs-podspec-properties-volumes.md "https://github.com/kitproj/kit/internal/types/pod#/$defs/PodSpec/properties/volumes")                                             |
| [semaphores](#semaphores)                                       | `object`  | Optional | cannot be null | [Untitled schema](pod-defs-podspec-properties-semaphores.md "https://github.com/kitproj/kit/internal/types/pod#/$defs/PodSpec/properties/semaphores")                                       |
| [env](#env)                                                     | `array`   | Optional | cannot be null | [Untitled schema](pod-defs-envvars.md "https://github.com/kitproj/kit/internal/types/pod#/$defs/PodSpec/properties/env")                                                                    |
| [envfile](#envfile)                                             | `array`   | Optional | cannot be null | [Untitled schema](pod-defs-envfile.md "https://github.com/kitproj/kit/internal/types/pod#/$defs/PodSpec/properties/envfile")                                                                |

## logLevel

LogLevel the log level to use by default.

`logLevel`

*   is optional

*   Type: `string` ([logLevel](pod-defs-podspec-properties-loglevel.md))

*   cannot be null

*   defined in: [Untitled schema](pod-defs-podspec-properties-loglevel.md "https://github.com/kitproj/kit/internal/types/pod#/$defs/PodSpec/properties/logLevel")

### logLevel Type

`string` ([logLevel](pod-defs-podspec-properties-loglevel.md))

## terminationGracePeriodSeconds

TerminationGracePeriodSeconds is the grace period for terminating the pod.

`terminationGracePeriodSeconds`

*   is optional

*   Type: `integer` ([terminationGracePeriodSeconds](pod-defs-podspec-properties-terminationgraceperiodseconds.md))

*   cannot be null

*   defined in: [Untitled schema](pod-defs-podspec-properties-terminationgraceperiodseconds.md "https://github.com/kitproj/kit/internal/types/pod#/$defs/PodSpec/properties/terminationGracePeriodSeconds")

### terminationGracePeriodSeconds Type

`integer` ([terminationGracePeriodSeconds](pod-defs-podspec-properties-terminationgraceperiodseconds.md))

## tasks

Tasks is a list of tasks that should be run.

`tasks`

*   is optional

*   Type: `object[]` ([Task](pod-defs-task.md))

*   cannot be null

*   defined in: [Untitled schema](pod-defs-tasks.md "https://github.com/kitproj/kit/internal/types/pod#/$defs/PodSpec/properties/tasks")

### tasks Type

`object[]` ([Task](pod-defs-task.md))

## volumes

Volumes is a list of volumes that can be mounted by containers belonging to the pod.

`volumes`

*   is optional

*   Type: `object[]` ([Volume](pod-defs-volume.md))

*   cannot be null

*   defined in: [Untitled schema](pod-defs-podspec-properties-volumes.md "https://github.com/kitproj/kit/internal/types/pod#/$defs/PodSpec/properties/volumes")

### volumes Type

`object[]` ([Volume](pod-defs-volume.md))

## semaphores

Semaphores is a list of semaphores that can be acquired by tasks.

`semaphores`

*   is optional

*   Type: `object` ([semaphores](pod-defs-podspec-properties-semaphores.md))

*   cannot be null

*   defined in: [Untitled schema](pod-defs-podspec-properties-semaphores.md "https://github.com/kitproj/kit/internal/types/pod#/$defs/PodSpec/properties/semaphores")

### semaphores Type

`object` ([semaphores](pod-defs-podspec-properties-semaphores.md))

## env

A list of environment variables.

`env`

*   is optional

*   Type: `object[]` ([EnvVar](pod-defs-envvar.md))

*   cannot be null

*   defined in: [Untitled schema](pod-defs-envvars.md "https://github.com/kitproj/kit/internal/types/pod#/$defs/PodSpec/properties/env")

### env Type

`object[]` ([EnvVar](pod-defs-envvar.md))

## envfile

Environment file (e.g. .env) to use

`envfile`

*   is optional

*   Type: `string[]`

*   cannot be null

*   defined in: [Untitled schema](pod-defs-envfile.md "https://github.com/kitproj/kit/internal/types/pod#/$defs/PodSpec/properties/envfile")

### envfile Type

`string[]`
