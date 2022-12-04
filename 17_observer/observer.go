package main

import "time"

type observer interface {
	update(g generator)
}

func sleep() {
	time.Sleep(time.Millisecond * 100)
}
