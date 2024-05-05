package task

type TaskState string

const (
	// task sent from client to manager waiting for being scheduled
	Pending TaskState = "PENDING"

	// scheduler decided the worker to host the task on, and sends it to the manager so the task is scheduled and ready to be hosted on this machine
	// or scheduler selected the task and task is hosted on the machine and worker is working on getting this task ready and running
	Scheduled TaskState = "SCHEDULED"

	// task is on the worker, and started running
	Running TaskState = "RUNNING"

	// task either get stopped by the client, or get it's work done successfully
	Completed TaskState = "COMPLETED"

	// task crashed while trying to stop it, or crashed while trying to be started, or crashed while trying to be scheduled
	Crashed TaskState = "CRASHED"
)
