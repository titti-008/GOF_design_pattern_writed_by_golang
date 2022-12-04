package main

import (
	"strings"
)

type bigString struct {
	bigchars []*bigChar
}

func newBigString(str string) (*bigString, error) {
	bs := new(bigString)
	f := getFactory()

	strArr := strings.Split(str, "")

	for _, char := range strArr {

		bc, err := f.getBigChar(char)
		if err != nil {
			return nil, err
		}

		bs.bigchars = append(bs.bigchars, bc)
	}
	return bs, nil
}

func (bs *bigString) print() {
	for _, bc := range bs.bigchars {
		bc.print()
	}
}
