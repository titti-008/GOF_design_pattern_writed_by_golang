package main

import (
	"fmt"
)

type listVisitor struct {
	currentDir string
}

func newListVisitor() *listVisitor {
	return &listVisitor{currentDir: ""}
}

var _ visitor = new(listVisitor)

func (l listVisitor) visitFile(f file) {
	fmt.Println(l.currentDir, "/", f)
}

func (l *listVisitor) visitDir(d *dir) {
	fmt.Println(l.currentDir, "/", d)
	saveDir := l.currentDir
	l.currentDir = l.currentDir + "/" + d.getName()
	for _, e := range d.entries {
		e.accept(l)
	}
	l.currentDir = saveDir
}
