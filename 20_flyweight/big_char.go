package main

import (
	"bufio"
	"fmt"
	"os"
)

type bigChar struct {
	charname string
	fontdata string
}

func newBigChar(charname string) (*bigChar, error) {
	c := new(bigChar)
	filename := "./fonts/big" + charname + ".txt"
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		c.fontdata += scanner.Text()
		c.fontdata += "\n"
	}

	fmt.Println(c.fontdata)

	return c, nil
}

func (c *bigChar) print() {
	fmt.Println(c.fontdata)
}
