package main

import (
	"testing"
)

func TestProbStNextHand(t *testing.T) {
	ps := NewProbStrategy()
	hand := ps.nextHand()
	if hand.handValue != 0 && hand.handValue != 1 && hand.handValue != 2 {
		t.Errorf("expected: 0~2, but got: %v", hand.handValue)
	}
}

func TestProbStGetSum(t *testing.T) {
	ps := NewProbStrategy()
	ps.history[1][0] = 3
	ps.history[1][1] = 4
	ps.history[1][2] = 5
	got := ps.getSum(0)
	expected := 3
	if got != expected {
		t.Errorf("expected: %v, but got: %v", expected, got)
	}
	got = ps.getSum(1)
	expected = 12
	if got != expected {
		t.Errorf("expected: %v, but got: %v", expected, got)
	}
}

func TestProbStStudy(t *testing.T) {
	ps := NewProbStrategy()
	ps.prevHandValue = ROCK
	ps.currentHandValue = ROCK
	ps.study(true)
	got := ps.history[ROCK][ROCK]
	expected := 2
	if got != expected {
		t.Errorf("expected: %v, but got: %v", expected, got)
	}

	ps.study(false)
	got = ps.history[ROCK][PAPER]
	expected = 2
	if got != expected {
		t.Errorf("expected: %v, but got: %v", expected, got)
	}
	got = ps.history[ROCK][SCISSORS]
	expected = 2
	if got != expected {
		t.Errorf("expected: %v, but got: %v", expected, got)
	}
}
