package main

import (
	"fmt"
)

type entry interface {
	getName() string
	getSize() int
	String() string
	accept(v visitor)
}

func toString(e entry) string {
	return fmt.Sprintf("%v (%v)", e.getName(), e.getSize())
}
