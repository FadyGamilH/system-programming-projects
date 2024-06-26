package worker

/*
- Node is the physical machine, worker is a node, manager is a node and uses many nodes (workers) to run tasks on them
- A node will have 0 or more tasks
*/
type Node struct {
	Name                    string // node-A
	IP                      string
	Role                    NodeRole
	MaxMemoryForTaskToUse   int
	MaxNumCoresForTaskToUse int
	TaskCount               int
	CurrentAllocatedMemory  int
	CurrentAllocatedDisk    int
}

type NodeRole string

const (
	WorkerNode  NodeRole = "worker"
	ManagerNode NodeRole = "manager"
)
