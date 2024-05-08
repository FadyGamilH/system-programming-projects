package worker

/*
Node is the physical machine, worker is a node, manager is a node and uses many nodes (workers) to run tasks on them
*/
type Node struct {
	Name                    string // node-A
	IP                      string
	Role                    string // is this node a manager or a worker (maybe we could define an enum later)
	MaxMemoryForTaskToUse   int
	MaxNumCoresForTaskToUse int
	TaskCount               int
	CurrentAllocatedMemory  int
	CurrentAllocatedTask    int
}
