package grid_cli

import (
	"fmt"

	"github.com/spf13/afero"
)

// Kernel struct with the syscall tree root and file system abstraction
type Kernel struct {
	root    *SyscallNode
	fs      afero.Fs
	modules map[string]Module // Known modules
}

// NewKernel initializes a new Kernel instance with embedded modules
func NewKernel() *Kernel {
	return &Kernel{
		root: &SyscallNode{
			Children: make(map[string]*SyscallNode),
			Modules:  []Module{},
		},
		fs:      afero.NewOsFs(),
		modules: make(map[string]Module),
	}
}

func (k *Kernel) addSyscall(parms ...interface{}) {
	current := k.root
	for _, parm := range parms {
		key := fmt.Sprintf("%v", parm)
		if _, exists := current.Children[key]; !exists {
			current.Children[key] = &SyscallNode{
				Children: make(map[string]*SyscallNode),
				Modules:  []Module{},
			}
		}
		current = current.Children[key]
	}
	// Assuming module is pre-initialized and available in context
	if module, exists := k.modules[fmt.Sprintf("%v", parms[len(parms)-1])]; exists {
		current.Modules = append(current.Modules, module)
	}
}

func (k *Kernel) findBestMatch(parms ...interface{}) *SyscallNode {
	current := k.root
	for _, parm := range parms {
		key := fmt.Sprintf("%v", parm)
		if next, exists := current.Children[key]; exists {
			current = next
		} else {
			break
		}
	}
	return current
}
