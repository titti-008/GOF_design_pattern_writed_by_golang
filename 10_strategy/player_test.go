package main

import (
	"testing"
)

func samplePlayer() *player {
	ws := newWinningStrategy()
	ws.prevHand = getHand(ROCK)
	ws.won = true
	p := newPlayer("player1", ws)
	return p
}

func TestPlayerNextHand(t *testing.T) {
	p := samplePlayer()
	if got, expected := p.nextHand().handValue, ROCK; got != expected {
		t.Errorf("expected: %v, but got: %v", expected, got)
	}
}

func TestWin(t *testing.T) {
	p := samplePlayer()
	p.win()
	if got, expected := p.winCount, 1; got != expected {
		t.Errorf("expected: %v, but got: %v", expected, got)
	}
	if got, expected := p.gameCount, 1; got != expected {
		t.Errorf("expected: %v, but got: %v", expected, got)
	}
}

func TestLose(t *testing.T) {
	p := samplePlayer()
	p.lose()
	if got, expected := p.loseCount, 1; got != expected {
		t.Errorf("expected: %v, but got: %v", expected, got)
	}
	if got, expected := p.gameCount, 1; got != expected {
		t.Errorf("expected: %v, but got: %v", expected, got)
	}
}

func TestEven(t *testing.T) {
	p := samplePlayer()
	p.even()
	if got, expected := p.gameCount, 1; got != expected {
		t.Errorf("expected: %v, but got: %v", expected, got)
	}
	if got, expected := p.loseCount, 0; got != expected {
		t.Errorf("expected: %v, but got: %v", expected, got)
	}
	if got, expected := p.winCount, 0; got != expected {
		t.Errorf("expected: %v, but got: %v", expected, got)
	}
}

func TestString(t *testing.T) {
	p := samplePlayer()
	p.gameCount = 10
	p.winCount = 2
	p.loseCount = 8
	if got, expected := p.String(), "[player1: 10 games, 2 win, 8 lose]"; got != expected {
		t.Errorf("expected: %v, but got: %v", expected, got)
	}

}
