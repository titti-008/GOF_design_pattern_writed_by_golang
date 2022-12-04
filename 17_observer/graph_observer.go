package main

import (
	"fmt"
)

type graphObserver struct{}

func (o graphObserver) update(g generator) {
	fmt.Print("GraphObserver: ")
	count := g.getNumber()
	for i := 0; i < count; i++ {
		fmt.Printf("%v", "*")
	}
	fmt.Print("\n")
	sleep()
}
