package main

import (
    "testing"
)

func TestDirectoryStringer(t *testing.T) {
    d := newDir("dirname")
    if got, expected := d.String(), "dirname (0)"; got != expected {
        t.Errorf("expected: %v, but got: %v", expected, got)
    }
}

func TestAdd(t *testing.T) {
    parent := newDir("parent")
    child := newDir("child")
    parent.add(child)

    if got, expected := parent.dir[0].getName(), "child"; got != expected {
        t.Errorf("expected: %v, but got: %v", expected, got)
    }

}
