package scheduler

type Scheduler interface {
	SelectCandidateWorkers()
	ScoreCandidateWorkers()
	PickWorker()
}
