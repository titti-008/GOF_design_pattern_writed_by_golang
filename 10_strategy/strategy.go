package main

type StrategyInterface interface {
	nextHand() *Hand
	study(win bool)
}
