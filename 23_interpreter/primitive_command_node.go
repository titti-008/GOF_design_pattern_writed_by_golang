package main

import (
	"errors"
	"fmt"
)

// <primitive command> ::= go | right | left
type primitiveCommandNode struct {
	name string
}

func newPrimitiveCommandNode() *primitiveCommandNode {
	return new(primitiveCommandNode)
}

func (n *primitiveCommandNode) parse(c *context) error {
	n.name = c.currentToken()
	if n.name == "" {
		return errors.New("Error: Missing <primitive command>")
	} else if n.name != "go" && n.name != "right" && n.name != "left" {
		return fmt.Errorf("Error: Unknown <primitive command>: '%v'", n.name)
	}
	c.skipToken(n.name)

	if n.name == "go" {
		c.window.walk()
	}

	if n.name == "right" {
		c.window.turnRight()
	}

	if n.name == "left" {
		c.window.turnLeft()
	}

	return nil
}

func (n *primitiveCommandNode) String() string {
	return n.name
}
