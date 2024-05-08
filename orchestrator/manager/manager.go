package manager

import (
	"log"
	"orchestrator/task"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

/*
Manager
  - Receives tasks from client and process them in FIFO basis
  - Has LocalStorage for tasks
  - Has LocalStorage for tasksEvents
*/
type Manager struct {
	TasksQueue             queue.Queue
	TasksLocalStorage      map[uuid.UUID]task.Task
	TaskEventsLocalStorage map[uuid.UUID]task.TaskEvent
	WorkersPool            []uuid.UUID
	// given a task name or id, find the worker that is running this task
	TaskWorkerMapByTaskName map[string]uuid.UUID
	TaskWorkerMapByTaskId   map[uuid.UUID]uuid.UUID

	// given worker id or name, find all tasks running on this worker
	WorkerTasksMapByWorkerId   map[uuid.UUID][]uuid.UUID
	WorkerTasksMapByWorkername map[string][]uuid.UUID
}

// after the manager talk to scheduler and know which worker will host this task, manager will talk to the worker directly to send the task
func (m *Manager) SendTaskToWorker() {
	log.Println("sending task to worker to be hosted")
}

// manager talks to scheduler and decide which worker
func (m *Manager) SpecifyWorkerToHostTask() {
	log.Println("specify worker to host the task")
}

// manager talks to wroker by calling CollectStatsPeriodicaly()
func (m *Manager) CollectTasksStats() {
	log.Println("collect stats of task from worker")
}

func (m *Manager) UpdateTask() {
	log.Println("call wroker to update a running task")
}
