package main

type janken int
const (
	ROCK     janken = 0
	SCISSORS janken = 1
	PAPER    janken = 2
)

type vs int

const (
	WIN  vs = 1
	LOSE vs = -1
	DRAW vs = 0
)

type Hand struct {
	name      string
	handValue janken
	hands     map[int]string
}

func newHand(handValue janken, name string) *Hand {
	h := new(Hand)
	h.handValue = handValue
	// h.hands = map[int]string{ROCK: "グー", SCISSORS: "チョキ", PAPER: "パー"}
	h.name = name
	return h
}

var hands []*Hand = []*Hand{
	newHand(ROCK, "グー"),
	newHand(SCISSORS, "チョキ"),
	newHand(PAPER, "パー"),
}

func getHand(handValue int) *Hand {
	return hands[handValue]

}

func (h *Hand) getValue(handValue int) int {
	return handValue
}

func (h *Hand) isStrongThan(refH *Hand) bool {
	return h.fight(refH) == WIN
}

func (h *Hand) fight(refH *Hand) vs {
	if h.handValue == refH.handValue {
		return DRAW
	} else if (h.handValue+1)%3 == refH.handValue {
		return WIN
	} else {
		return LOSE
	}
}
