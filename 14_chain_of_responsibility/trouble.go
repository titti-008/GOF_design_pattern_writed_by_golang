package main

import (
	"fmt"
)

type trouble struct {
	number int
}

func newTrouble(num int) *trouble {
	return &trouble{number: num}
}

func (t trouble) getNumber() int {
	return t.number
}

func (t trouble) String() string {
	return fmt.Sprintf("[Trouble %v]", t.number)
}
