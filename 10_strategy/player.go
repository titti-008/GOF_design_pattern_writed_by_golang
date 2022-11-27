package main

import (
	"fmt"
)

type player struct {
	name      string
	strategy  StrategyInterface
	winCount  int
	loseCount int
	gameCount int
}

func newPlayer(name string, s StrategyInterface) *player {
	p := new(player)
	p.name = name
	p.strategy = s
	return p
}

func (p *player) nextHand() *Hand {
	return p.strategy.nextHand()
}

func (p *player) win() {
	p.strategy.study(true)
	p.winCount++
	p.gameCount++
}

func (p *player) lose() {
	p.strategy.study(false)
	p.loseCount++
	p.gameCount++
}

func (p *player) even() {
	p.gameCount++
}

func (p player) String() string {
	return fmt.Sprintf("[%v: %v games, %v win, %v lose]", p.name, p.gameCount, p.winCount, p.loseCount)
}
