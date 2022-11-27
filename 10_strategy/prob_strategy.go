package main

import (
	"math/rand"
	"time"
)

type ProbStrategy struct {
	prevHandValue    int
	currentHandValue int
	history          [][]int
}

func newProbStrategy() *ProbStrategy {
	ps := new(ProbStrategy)
	ps.prevHandValue = 0
	ps.currentHandValue = 0
	ps.history = [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}
	return ps
}

func (ps *ProbStrategy) nextHand() *Hand {
	rand.Seed(time.Now().UnixNano())
	bet := rand.Intn(ps.getSum(ps.currentHandValue))
	handValue := 0
	if bet < ps.history[ps.currentHandValue][0] {
		handValue = 0
	} else if bet < ps.history[ps.currentHandValue][0]+ps.history[ps.currentHandValue][1] {
		handValue = 1
	} else {
		handValue = 2
	}
	ps.prevHandValue = ps.currentHandValue
	ps.currentHandValue = handValue
	return getHand(handValue)
}

func (ps *ProbStrategy) getSum(handValue int) int {
	sum := 0
	for i := range ps.history[handValue] {
		sum += ps.history[handValue][i]
	}
	return sum
}

func (ps *ProbStrategy) study(won bool) {
	if won {
		ps.history[ps.prevHandValue][ps.currentHandValue]++
	} else {
		ps.history[ps.prevHandValue][(ps.currentHandValue+1)%3]++
		ps.history[ps.prevHandValue][(ps.currentHandValue+2)%3]++
	}

}
