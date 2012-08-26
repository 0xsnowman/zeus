package zeusmaster

import "net"

type ProcessTree struct {
	Root *SlaveNode
	ExecCommand string
	slavesByName map[string]*SlaveNode
	commandsByName map[string]*CommandNode
}

type ProcessTreeNode struct {
	Name string
	Action string
}

type SlaveNode struct {
	ProcessTreeNode
	Socket *net.UnixConn
	Pid int
	Slaves []*SlaveNode
	Commands []*CommandNode
	Features map[string]bool
}

type CommandNode struct {
	ProcessTreeNode
	Aliases []string
}

func (tree *ProcessTree) NewCommandNode(name string, aliases []string) *CommandNode {
	x := &CommandNode{}
	x.Name = name
	tree.commandsByName[name] = x
	return x
}

func (tree *ProcessTree) NewSlaveNode(name string) *SlaveNode {
	x := &SlaveNode{}
	x.Name = name
	tree.slavesByName[name] = x
	return x
}

func (tree *ProcessTree) FindSlaveByName(name string) *SlaveNode {
	if name == "" {
		return tree.Root
	}
	return tree.slavesByName[name]
}

func (tree *ProcessTree) FindCommandByName(name string) *CommandNode {
	return tree.commandsByName[name]
}
