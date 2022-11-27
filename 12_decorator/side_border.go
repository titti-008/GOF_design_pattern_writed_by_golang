package main

type sideBorder struct {
	display    displayInterface
	borderChar string
}

var _ displayInterface = new(sideBorder)

func newSideBorder(d displayInterface, c string) *sideBorder {
	sb := new(sideBorder)
	sb.display = d
	sb.borderChar = c
	return sb
}

func (sb *sideBorder) getColumns() int {
	return 1 + sb.display.getColumns() + 1
}

func (sb *sideBorder) getRows() int {
	return sb.display.getRows()
}

func (sb *sideBorder) getRowText(row int) string {
	return sb.borderChar + sb.display.getRowText(row) + sb.borderChar
}
