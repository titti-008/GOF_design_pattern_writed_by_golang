package game

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type gamer struct {
	money      int
	fruits     []string
	fruitsName []string
}

func NewGamer(money int) *gamer {
	rand.Seed(time.Now().UnixNano())
	g := new(gamer)
	g.money = money
	g.fruitsName = []string{"りんご", "ぶどう", "バナナ", "みかん"}
	return g
}

func (g gamer) randBool() bool {
	return rand.Intn(2) == 1
}

func (g gamer) GetMoney() int {
	return g.money
}

func (g *gamer) Bet() {
	dice := rand.Intn(6) + 1
	if dice == 1 {
		g.money += 100
		fmt.Println("所持金が増えました")
	} else if dice == 2 {
		g.money /= 2
		fmt.Println("所持金が半分になりました")
	} else if dice == 6 {
		f := g.GetFruits()
		fmt.Println("フルーツ(", f, ")をもらいました。")
		g.fruits = append(g.fruits, f)
	} else {
		fmt.Println("何も起こりませんでした。")
	}
}

func (g gamer) CreateMemento() *memento {
	m := newMement(g.money)
	for _, f := range g.fruits {
		if strings.HasPrefix(f, "おいしい") {
			m.addFruit(f)
		}
	}
	return m
}

func (g *gamer) RestoreMemento(m *memento) {
	g.money = m.GetMoney()
	g.fruits = m.getFruits()

}

func (g *gamer) GetFruits() string {
	f := g.fruitsName[rand.Intn(len(g.fruitsName))]
	if g.randBool() {
		return "おいしい" + f
	}
	return f
}

func (g *gamer) String() string {
	return fmt.Sprintf("[money = %v, fruits = %v]", g.money, g.fruits)
}
