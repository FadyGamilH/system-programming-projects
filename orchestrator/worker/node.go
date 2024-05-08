package worker

/*
Node is the physical machine, worker is a node, manager is a node and uses many nodes (workers) to run tasks on them
*/
type Node struct {
	Name                    string // node-A
	IP                      string
	Role                    NodeRole
	MaxMemoryForTaskToUse   int
	MaxNumCoresForTaskToUse int
	TaskCount               int
	CurrentAllocatedMemory  int
	CurrentAllocatedTask    int
}

type NodeRole string

const (
	WorkerNode  NodeRole = "worker"
	ManagerNode NodeRole = "manager"
)
