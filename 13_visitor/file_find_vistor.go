package main

import (
	"strings"
)

type fileFindVisitor struct {
	id         string
	founds     []entry
	currentDir string
}

func newFFV(id string) *fileFindVisitor {
	return &fileFindVisitor{id: id}
}

var _ visitor = new(fileFindVisitor)

func (v *fileFindVisitor) visitFile(f file) {
	if v.isIncluded(f.getName()) {
		v.founds = append(v.founds, f)
	}
}

func (v *fileFindVisitor) visitDir(d *dir) {
	for _, e := range d.entries {
		e.accept(v)
	}
}

func (v *fileFindVisitor) isIncluded(name string) bool {
	return strings.Contains(name, v.id)
}

func (v *fileFindVisitor) getFounds() []entry {
	return v.founds
}
