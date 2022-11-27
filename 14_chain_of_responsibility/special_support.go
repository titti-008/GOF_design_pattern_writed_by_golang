package main

type specialSupport struct {
	number int
	support
}

var _ supportInterface = new(specialSupport)

func newSpecialSupport(name string, num int) *specialSupport {
	s := new(specialSupport)
	s.name = name
	s.number = num
	return s
}

func (s *specialSupport) resolve(t trouble) bool {
	return t.getNumber() == s.number
}

func (s *specialSupport) supportTrouble(t trouble) {
	supportTrouble(s, t)
}
