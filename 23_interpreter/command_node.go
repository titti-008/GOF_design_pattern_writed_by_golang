package main

type commandNode struct {
	node node
}

var _ node = new(commandNode)

func newCommandNode() *commandNode {
	return new(commandNode)
}

func (n *commandNode) parse(c *context) error {
	if c.currentToken() == "repeat" {
		n.node = newRepeatCommandNode()
	} else {
		n.node = newPrimitiveCommandNode()
	}
	err := n.node.parse(c)

	return err
}

func (n *commandNode) String() string {
	return n.node.String()
}
