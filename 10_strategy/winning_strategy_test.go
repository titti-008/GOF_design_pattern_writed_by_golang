package main

import (
	// "reflect"
	"testing"
)

func TestWinStNextHand(t *testing.T) {
	prevHand := getHand(ROCK)
	ws := newWinningStrategy()
	ws.prevHand = prevHand
	ws.won = true
	got := ws.nextHand()
	expected := prevHand
	if got != expected {
		t.Errorf("expected: %v, but got: %v", expected, got)
	}

}

func TestWinStStudy(t *testing.T) {
	ws := newWinningStrategy()
	ws.study(true)

	got := ws.won
	expected := true
	if got != expected {
		t.Errorf("expected: %v, but got: %v", expected, got)
	}

}
