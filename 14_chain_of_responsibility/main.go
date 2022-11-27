package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello chain of responsibirity")

	alice := newNoSupport("Alice")
	bob := newLimitSupport("Bob", 100)
	charlie := newSpecialSupport("Charlie", 429)
	diana := newLimitSupport("Diana", 200)
	elmo := newOddSupport("Elmo")
	fred := newLimitSupport("Fred", 300)

	alice.setNext(bob).setNext(charlie).setNext(diana).setNext(elmo).setNext(fred)

	for i := 0; i < 500; i += 33 {
		alice.supportTrouble(*newTrouble(i))
	}
}
