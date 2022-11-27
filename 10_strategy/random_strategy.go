package main

import (
    "time"
    "math/rand"
)
var _ StrategyInterface = &randomStrategy{}


type randomStrategy struct {
    
}

func (rs *randomStrategy) nextHand() *Hand {
    rand.Seed(time.Now().UnixNano())

    return getHand(rand.Intn(2))
}

func (rs *randomStrategy) study (won bool) {
}
