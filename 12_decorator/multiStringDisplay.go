package main

type multiStringDisplay struct {
	lines []string
}

var _ displayInterface = new(multiStringDisplay)

func newMultiStringDisplay() *multiStringDisplay {
	return new(multiStringDisplay)
}

func (d *multiStringDisplay) add(line string) {
	d.lines = append(d.lines, line)
}

func (d multiStringDisplay) getColumns() int {
	max := 0
	for _, line := range d.lines {
		l := len(line)
		if max < l {
			max = l
		}
	}
	return max
}

func (d multiStringDisplay) getRows() int {
	return len(d.lines)
}

func (d multiStringDisplay) getRowText(row int) string {
	result := d.lines[row]
	diff := d.getColumns() - len(result)
	for i := 0; i < diff; i++ {
		result += " "
	}

	return result
}
