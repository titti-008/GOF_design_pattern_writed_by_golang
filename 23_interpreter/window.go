package main

import (
	"errors"
	"fmt"
)

type position struct {
	X int
	Y int
}

var right = position{X: 1, Y: 0}
var left = position{X: -1, Y: 0}
var up = position{X: 0, Y: -1}
var down = position{X: 0, Y: 1}

type window struct {
	board     [][]string
	length    int
	start     *position
	now       *position
	finished  *position
	direction position
}

func newWindow(length int, start position) *window {
	w := new(window)
	w.length = length
	w.board = make([][]string, length)
	for y := 0; y < length; y++ {
		w.board[y] = make([]string, length)
		for x := 0; x < length; x++ {
			w.board[y][x] = "."
		}
	}

	w.board[start.Y][start.X] = "S"
	w.start = &start
	w.now = &position{X: w.start.X, Y: w.start.Y}
	w.direction = up

	return w
}

func (w *window) walk() error {
	x := w.now.X + w.direction.X
	y := w.now.Y + w.direction.Y
	if (0 <= x && x < w.length) && (0 <= y && y < w.length) {
		w.now.X = x
		w.now.Y = y
		w.board[y][x] = "*"

	} else {
		return errors.New("そのpositionには移動できません。")
	}

	return nil
}

func (w *window) turnRight() {
	x := w.direction.X
	y := w.direction.Y
	w.direction.X = -y
	w.direction.Y = x
}

func (w *window) turnLeft() {
	x := w.direction.X
	y := w.direction.Y
	w.direction.X = y
	w.direction.Y = -x
}

func (w window) print() {
	result := "board:\n"
	for y := range w.board {
		for x := range w.board[y] {
			result += w.board[y][x]
		}
		result += "\n"
	}

	fmt.Println(result)
}
