package main

import (
	"fmt"
)

type plusObserver struct{}

func (o *plusObserver) update(g generator) {

	fmt.Print("plusObserver: ")
	count := g.getNumber()
	result := "\n"
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			result += "+"
		}
		result += "\n"
	}
	fmt.Print(result)
	sleep()
}
