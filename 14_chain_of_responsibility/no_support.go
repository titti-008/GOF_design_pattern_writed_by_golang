package main

type noSupport struct {
	support
}

var _ supportInterface = new(noSupport)

func newNoSupport(name string) *noSupport {
	s := new(noSupport)
	s.name = name
	return s
}

func (s *noSupport) resolve(t trouble) bool {
	return false
}

func (s *noSupport) supportTrouble(t trouble) {
	supportTrouble(s, t)
}
