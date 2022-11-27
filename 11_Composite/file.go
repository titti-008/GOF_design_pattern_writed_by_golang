package main

import (
    "fmt"
)

type file struct{
    name string
    size int
}

var _ entryInterface = &file{}

func newFile(name string, size int) *file {
    f := new(file)
    f.name = name
    f.size = size
    return f
}

func (f *file) getName() string {
    return f.name
}

func (f *file) getSize() int {
    return f.size
}

func (f *file) printList(prefix string) {
    fmt.Println(prefix, "/", f)
}

func (f *file) String() string {
    return fmt.Sprintf("%v (%v)", f.getName(), f.getSize())
}
