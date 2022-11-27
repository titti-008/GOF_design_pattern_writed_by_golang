package main

type oddSupport struct {
	support
}

func newOddSupport(name string) *oddSupport {
	s := new(oddSupport)
	s.name = name
	return s
}

func (s *oddSupport) resolve(t trouble) bool {
	return t.getNumber()%2 == 1
}

func (s *oddSupport) supportTrouble(t trouble) {
	supportTrouble(s, t)
}
