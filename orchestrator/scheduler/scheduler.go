package scheduler

/*
The scheduler is the way of knowing where to host the task, its simply an algorithm and we will build two types of it
  - Round Robin
  - High Availability
*/
type Scheduler interface {
	SelectCandidateWorkers()
	ScoreCandidateWorkers()
	PickWorker()
}
