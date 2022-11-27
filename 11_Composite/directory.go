package main

import (
    "fmt"
)

type directory struct {
    dir []entryInterface
    name string
}

var _ entryInterface = new(directory)

func newDir (name string) *directory{
    d := new(directory)
    d.name = name

    return d
}

func (d *directory) getName ()string {
    return d.name
}

func (d *directory) getSize () int {
    result := 0
    for _, entry := range d.dir {
        result += entry.getSize()
    }
    return result
}

func (d *directory) printList(prefix string) {
    fmt.Println(prefix, d)
    for _, entry := range d.dir {
        p := fmt.Sprintf("%v/%v", prefix, d.name)
        entry.printList(p)
    }
}

func (d *directory) String() string {
    return fmt.Sprintf("%v (%v)", d.getName(), d.getSize())
}

func (d *directory) add (entry entryInterface) {
    d.dir = append(d.dir, entry)
}
