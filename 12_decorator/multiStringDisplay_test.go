package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	d := newMultiStringDisplay()
	d.add("line1")
	if got, expected := d.lines[0], "line1"; got != expected {
		t.Errorf("expected: %v, but got: %v", expected, got)
	}
	d.add("line2")
	if got, expected := d.lines[1], "line2"; got != expected {
		t.Errorf("expected: %v, but got: %v", expected, got)
	}

	d.add("123456789")
	if got, expected := d.getColumns(), 9; got != expected {
		t.Errorf("expected: %v, but got: %v", expected, got)
	}

	if got, expected := d.getRows(), 3; got != expected {
		t.Errorf("expected: %v, but got: %v", expected, got)
	}

	if got, expected := d.getRowText(0), "line1"; got != expected {
		t.Errorf("expected: %v, but got: %v", expected, got)
	}
}

func TestGetRows(t *testing.T) {

}
