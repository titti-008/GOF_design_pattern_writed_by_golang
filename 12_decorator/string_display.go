package main

import (
	"errors"
)

type stringDisplay struct {
	str string
}

var _ displayInterface = new(stringDisplay)

func newStringDisplay(str string) *stringDisplay {
	return &stringDisplay{str: str}
}

func (sd stringDisplay) getColumns() int {
	return len(sd.str)

}

func (sd stringDisplay) getRows() int {
	return 1
}

func (sd stringDisplay) getRowText(row int) string {
	if row != 0 {
		panic(errors.New("index out of bounds exception"))
	}
	return sd.str
}
