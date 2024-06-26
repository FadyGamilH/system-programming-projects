package task

import "github.com/docker/go-connections/nat"

/*
Config struct represents the configs required for a task to be running on docker successfully
  - Name will represent both the name of the task and the container (when task becoming a running contaner)
  - RestartPolicy is the instruction on what should be done when the docker container stopped or failed ?
*/
type DockerConfig struct {
	ExposedPorts nat.PortSet
	Env          []string
	CMD          []string
	// will be used by the scheduler, these are the actuall required memory and disk
	// the scheduler will use them to find the wroker that will be able to hist a running task with these specs
	Memory        int64
	Disk          int64
	Name          string // describes the name of the task and the container (both)
	RestartPolicy RestartPolicy
	Image         string
	CPU           float64
	AttachStdin   bool
	AttachStdout  bool
	AttachStderr  bool
}

type RestartPolicy string

const (
	Empty         = ""               // don't restart the container if it failed (default)
	OnFailure     = "ON_FAILURE"     // when a fail reason is a non zero error
	Always        = "ALWAYS"         // once docker container stopped, try to restart it
	UnlessStopped = "UNLESS_STOPPED" // if the container stopped via the docker stop command (intentionally)
)
