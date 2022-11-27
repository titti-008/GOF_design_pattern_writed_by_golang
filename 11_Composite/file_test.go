package main

import (
    "testing"
)

func TestFileStringer(t *testing.T) {
    f := newFile("filename", 1111)
    if got, expected := f.String(), "filename (1111)"; got != expected {
        t.Errorf("expected: %v, but got: %v", expected, got)
    }
}
