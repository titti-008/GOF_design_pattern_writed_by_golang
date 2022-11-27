package main

import (
	"fmt"
)

type displayInterface interface {
	getColumns() int
	getRows() int
	getRowText(row int) string
}

func show(d displayInterface) {
	result := ""
	for i := 0; i < d.getRows(); i++ {
		result += d.getRowText(i)
		result += "\n"
	}
	fmt.Println(result)
}
