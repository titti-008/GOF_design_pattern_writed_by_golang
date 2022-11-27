package main

import (
	"fmt"
)

type supportInterface interface {
	setNext(next supportInterface) supportInterface
	getNext() supportInterface
	supportTrouble(trouble)
	String() string
	resolve(trouble) bool
	done(trouble)
	fail(trouble)
}

type support struct {
	name string
	next supportInterface
}

func supportTrouble(s supportInterface, t trouble) {
	for obj := s; true; obj = obj.getNext() {
		if obj.resolve(t) {
			obj.done(t)
			break
		} else if obj.getNext() == nil {
			obj.fail(t)
			break
		}
	}
}

func (s *support) setNext(next supportInterface) supportInterface {
	s.next = next
	return s.next
}

func (s *support) String() string {
	return fmt.Sprintf("[%v]", s.name)
}

func (s *support) getNext() supportInterface {
	return s.next
}

func (s *support) done(t trouble) {
	fmt.Printf("%v is resolved by %v.\n", t, s)
}

func (s *support) fail(t trouble) {
	fmt.Printf("%v cannot be resolved.\n", t)
}
