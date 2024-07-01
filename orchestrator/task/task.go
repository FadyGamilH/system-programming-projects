package task

import (
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
)

type Task struct {
	Id            uuid.UUID // 128 bits
	Name          string
	State         TaskState   // at the end its a string
	DockerImage   string      // each task will be a program running on a docker container and this container is from a specific image
	RestartPolicy string      // how our orchasterator will behave if the task stopped or crashed
	ExposedPorts  nat.PortSet // map of port to struct, where the port is a string
	PortsBinding  map[string]string
	Disk          int
	Memory        int
	StartTime     time.Time
	EndTime       time.Time
}

// we want an event to reprsent the transition of a task from state-A to state-B (specially this is used when user need to stop or start the task)
type TaskEvent struct {
	Task                Task
	Id                  uuid.UUID
	StateToTransitionTo TaskState
	RequestedAt         time.Time
}
