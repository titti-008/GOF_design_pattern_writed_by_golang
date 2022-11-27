package main

type file struct {
	name string
	size int
}

// var _ element = new(file)
var _ entry = new(file)

func newFile(name string, size int) *file {
	return &file{name: name, size: size}
}

func (f file) getName() string {
	return f.name
}

func (f file) getSize() int {
	return f.size
}

func (f file) String() string {
	return toString(f)
}

func (f file) accept(v visitor) {
	v.visitFile(f)
}
