package main

import (
	"fmt"
)

type fullBorder struct {
	display displayInterface
}

func newFullBorder(d displayInterface) *fullBorder {
	return &fullBorder{display: d}
}

var _ displayInterface = new(fullBorder)

func (fb fullBorder) getColumns() int {
	return 1 + fb.display.getColumns() + 1
}
func (fb fullBorder) getRows() int {
	return 1 + fb.display.getRows() + 1
}

func (fb fullBorder) getRowText(row int) string {
	if row == 0 {
		return fmt.Sprintf("+%v+", fb.makeLine("-", fb.display.getColumns()))
	} else if row == fb.display.getRows()+1 {
		return fmt.Sprintf("+%v+", fb.makeLine("-", fb.display.getColumns()))
	}
	return fmt.Sprintf("|%v|", fb.display.getRowText(row-1))
}

func (fb fullBorder) makeLine(ch string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += ch
	}
	return result
}
