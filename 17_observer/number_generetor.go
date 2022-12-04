package main

type generator interface {
	addObserver(o observer)
	deleteObserver(o observer)
	notifyObserver()
	getNumber() int
	execute()
}

type numberGenerator struct {
	observers []observer
}

func addObserver(g *numberGenerator, o observer) {
	g.observers = append(g.observers, o)
}

func deleteObserver(observers []observer, o observer) []observer {
	var removed []observer
	for _, observer := range observers {
		if observer != o {
			removed = append(observers, o)
		}
	}

	return removed
}
