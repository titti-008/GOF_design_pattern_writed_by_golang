package main

import (
	"fmt"
)

type digitObserver struct {
}

var _ observer = new(digitObserver)

func (o digitObserver) update(g generator) {
	fmt.Println("digitObserver: ", g.getNumber())
	sleep()
}
