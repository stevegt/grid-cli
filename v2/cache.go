package grid_cli

// SyscallNode represents a node in the hierarchical syscall tree
type SyscallNode struct {
	Modules  []Module
	Children map[string]*SyscallNode
}
