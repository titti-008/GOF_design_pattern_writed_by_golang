package main

func main() {
	// g := newRandNumGen()
	g := newincreNumGen(10, 50, 5)
	dObserver := new(digitObserver)
	gObserver := new(graphObserver)
	pObserver := new(plusObserver)
	g.addObserver(dObserver)
	g.addObserver(gObserver)
	g.addObserver(pObserver)
	g.execute()
}
