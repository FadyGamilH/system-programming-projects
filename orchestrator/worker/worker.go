package worker

import (
	"orchestrator/task"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

/*
Worker Features :

  - Host and Run the task as a docker container
  - Provides stats to the manager
  - keep track of the stats of its hosted tasks
  - Has a local storage
*/
type Worker struct {
	TasksQueue queue.Queue
	Storage    map[uuid.UUID]*task.Task
	Name       string
	TasksCount int64
}

func (w *Worker) CollectStatsPeriodicaly() error {
	return nil
}

func (w *Worker) StartTask() error {
	return nil
}

func (w *Worker) StopTask() error {
	return nil
}

/*
ProcessTask() method process the task by performing one of two options:
  - specify the task state
  - based on the task state it will stop or start the task
*/
func (w *Worker) ProcessTask() error {
	return nil
}
