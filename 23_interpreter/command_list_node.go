package main

import (
	"errors"
)

// <command list> ::= <command>* end
type commandListNode struct {
	list []node
}

var _ node = new(commandListNode)

func newCommandListNode() *commandListNode {
	n := new(commandListNode)
	n.list = []node{}
	return n
}

func (n *commandListNode) parse(c *context) error {
	for {
		if c.currentToken() == "" {
			return errors.New("Error: Missing 'end'")
		} else if c.currentToken() == "end" {
			c.skipToken("end")
			break
		} else {
			commandNode := newCommandNode()
			err := commandNode.parse(c)
			if err != nil {
				return err
			}
			n.list = append(n.list, commandNode)
		}
	}

	return nil
}

func (n *commandListNode) String() string {
	result := ""
	for _, node := range n.list {
		result += node.String()
	}
	return result
}
