package main

type sizeVisitor struct {
	size int
}

var _ visitor = new(sizeVisitor)

func (v *sizeVisitor) visitFile(f file) {
	v.size += f.getSize()
}

func (v *sizeVisitor) visitDir(d *dir) {
	for _, e := range d.entries {
		e.accept(v)
	}
}

func (v *sizeVisitor) getSize() int {
	return v.size
}
