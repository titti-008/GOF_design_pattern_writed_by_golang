package main

import (
	"fmt"
	"github.com/titti-008/design_pattern/18_memento/game"
	"time"
)

func main() {
	money := 100
	g := game.NewGamer(money)
	m := g.CreateMemento()
	for i := 0; i < money; i++ {
		fmt.Println("====", i)
		fmt.Println("現状: ", g)

		g.Bet()

		fmt.Println("所持金は", g.GetMoney(), "円になりました。")

		if g.GetMoney() > m.GetMoney() {
			fmt.Println("細部増えたので現在の状態を保存しておこう!")
			m = g.CreateMemento()
		} else if g.GetMoney() < m.GetMoney() {
			fmt.Println("だいぶ減ったので以前の状態を復元しよう!")
			g.RestoreMemento(m)
		}
		time.Sleep(time.Second)
		fmt.Println("")
	}
}
