package main

import (
	"testing"
)

func TestInclude(t *testing.T) {
	v := new(fileFindVisitor)
	v.id = "txt"
	if got, expected := v.isIncluded("file.txt"), true; got != expected {
		t.Errorf("expected: %v, but got: %v", expected, got)
	}
	if got, expected := v.isIncluded("file.doc"), false; got != expected {
		t.Errorf("expected: %v, but got: %v", expected, got)
	}
}
