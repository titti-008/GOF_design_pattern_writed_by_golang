package main

type increNumGen struct {
	numberGenerator
	number, max, diff int
}

var _ generator = new(increNumGen)

func newincreNumGen(min, max, diff int) *increNumGen {
	g := &increNumGen{max: max, diff: diff}
	g.number = min

	return g
}

func (g *increNumGen) addObserver(o observer) {
	g.observers = append(g.observers, o)
}

func (g *increNumGen) deleteObserver(o observer) {
	g.observers = deleteObserver(g.observers, o)
}

func (g *increNumGen) getNumber() int {
	return g.number
}

func (g *increNumGen) execute() {
	for {
		g.notifyObserver()
		g.number += g.diff
		if g.number >= g.max {
			break
		}
	}
}

func (g *increNumGen) notifyObserver() {
	for _, o := range g.observers {
		o.update(g)
	}
}
