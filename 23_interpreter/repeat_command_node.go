package main

import (
	"fmt"
)

type repeatCommandNode struct {
	number          int
	commandListNode node
}

func newRepeatCommandNode() *repeatCommandNode {
	return &repeatCommandNode{}
}

func (n *repeatCommandNode) parse(c *context) error {
	c.skipToken("repeat")
	num, err := c.currentNumber()
	if err != nil {
		return err
	}
	n.number = num

	c.nextToken()
	n.commandListNode = newCommandListNode()

	err = n.parse(c)
	return err
}

func (n *repeatCommandNode) String() string {
	return fmt.Sprintf("[repeat %d %s]", n.number, n.commandListNode)
}
