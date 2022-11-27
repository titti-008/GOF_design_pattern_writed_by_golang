package main

import (
	"math/rand"
	"time"
)

type WinningStrategy struct {
	won      bool
	prevHand *Hand
}

func newWinningStrategy() *WinningStrategy {
	ws := new(WinningStrategy)
	ws.won = false
	return ws
}

func (ws *WinningStrategy) nextHand() *Hand {
	rand.Seed(time.Now().UnixNano())
	if ws.won {
		return ws.prevHand
	}

    ws.prevHand = getHand(rand.Intn(2))
	return ws.prevHand
}

func (ws *WinningStrategy) study(win bool) {
	ws.won = win
}
