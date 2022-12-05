package main

// <program> ::= program <command list>
type programNode struct {
	commandListNode node
}

var _ node = new(programNode)

func (n *programNode) parse(c *context) error {
	c.skipToken("program")
	n.commandListNode = newCommandListNode()
	err := n.commandListNode.parse(c)
	if err != nil {
		return err
	}

	return nil
}

func (n *programNode) String() string {
	return "[program " + n.commandListNode.String() + "]"
}
