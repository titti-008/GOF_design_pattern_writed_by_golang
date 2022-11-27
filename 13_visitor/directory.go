package main

type dir struct {
	name    string
	entries []entry
	size    int
}

var _ entry = new(dir)

func newDir(name string) *dir {
	return &dir{name: name}
}

func (d dir) getName() string {
	return d.name
}

func (d dir) getSize() int {
	v := new(sizeVisitor)
	d.accept(v)
	return v.getSize()
}

func (d dir) String() string {
	return toString(d)
}

func (d dir) accept(v visitor) {
	v.visitDir(&d)
}

func (d *dir) add(e entry) {
	d.entries = append(d.entries, e)
}
