package main

import (
	"testing"
)

func TestGetHand(t *testing.T) {
	h := getHand(ROCK)
	got := h.name
	expected := "グー"
	if got != expected {
		t.Errorf("expected: %v, but got: %v", expected, got)
	}

	// if h.hands[1] != "チョキ" {
	// 	t.Errorf("expected: %v, but got: %v", "チョキ", h.hands[1])
	// }

}

func TestGetValue(t *testing.T) {
	h := getHand(PAPER)
	got := h.getValue(2)
	expected := 2

	if got != expected {
		t.Errorf("expected: %v, but got: %v", expected, got)
	}
}

func TestIsStrongThan(t *testing.T) {
	h0 := getHand(ROCK)
	h1 := getHand(PAPER)
	got := h1.isStrongThan(h0)
	expected := true
	if got != expected {
		t.Errorf("expected: %v, but got: %v", expected, got)
	}

	got = h0.isStrongThan(h1)
	expected = false
	if got != expected {
		t.Errorf("expected: %v, but got: %v", expected, got)
	}
}

func TestFight(t *testing.T) {
	h0 := getHand(ROCK)
	h1 := getHand(PAPER)
	got := h1.fight(h0)
	expected := 1
	if got != expected {
		t.Errorf("expected: %v, but got: %v", expected, got)
	}

	got = h0.fight(h1)
	expected = -1
	if got != expected {
		t.Errorf("expected: %v, but got: %v", expected, got)
	}

	got = h0.fight(h0)
	expected = 0
	if got != expected {
		t.Errorf("expected: %v, but got: %v", expected, got)
	}
}
