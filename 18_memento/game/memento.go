package game

type memento struct {
	money  int
	fruits []string
}

func newMement(money int) *memento {
	m := new(memento)
	m.money = money
	return m
}

func (m *memento) GetMoney() int {
	return m.money
}

func (m *memento) addFruit(fruit string) {
	m.fruits = append(m.fruits, fruit)
}

func (m *memento) getFruits() []string {
	return m.fruits
}
