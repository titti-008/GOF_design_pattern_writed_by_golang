package main

import (
	"math/rand"
	"time"
)

type randNumGen struct {
	numberGenerator
	number int
}

var _ generator = new(randNumGen)

func newRandNumGen() *randNumGen {
	g := new(randNumGen)
	return g
}

func (g *randNumGen) addObserver(o observer) {
	g.observers = append(g.observers, o)
}

func (g *randNumGen) deleteObserver(o observer) {
	g.observers = deleteObserver(g.observers, o)
}

func (g *randNumGen) getNumber() int {
	return g.number
}

func (g *randNumGen) execute() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		g.number = rand.Intn(50)
		g.notifyObserver()
	}
}

func (g *randNumGen) notifyObserver() {
	for _, o := range g.observers {
		o.update(g)
	}
}
