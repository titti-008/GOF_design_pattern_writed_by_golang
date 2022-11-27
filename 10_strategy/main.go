package main

import (
    "log"
)

func main() {
	log.Println("hello strategy")
    // p1 := newPlayer("player1", newWinningStrategy())
    p1 := newPlayer("player1", new(randomStrategy))
    p2 := newPlayer("player2", newProbStrategy())
    for i:=0; i<10000; i++{
        nextHand1 := p1.nextHand()
        nextHand2 := p2.nextHand()
        if nextHand1.isStrongThan(nextHand2){
            log.Println("Winner: ", p1)
            p1.win()
            p2.lose()
        } else if nextHand2.isStrongThan(nextHand1){
            log.Println("Winner: ", p2)
            p2.win()
            p1.lose()
        } else {
            log.Println("even..")
            p1.even()
            p2.even()
        }
    }

    log.Println("Total result: ")
    log.Println(p1)
    log.Println(p2)
}
