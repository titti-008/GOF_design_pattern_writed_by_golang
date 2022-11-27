package main

type limitSupport struct {
	support
	limit int
}

var _ supportInterface = new(limitSupport)

func newLimitSupport(name string, limit int) *limitSupport {
	s := new(limitSupport)
	s.name = name
	s.limit = limit
	return s
}

func (s *limitSupport) resolve(t trouble) bool {
	return t.getNumber() < s.limit
}

func (s *limitSupport) supportTrouble(t trouble) {
	supportTrouble(s, t)
}
