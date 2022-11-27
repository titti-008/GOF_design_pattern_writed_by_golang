package main

import (
	"github.com/titti-008/design_pattern/15_facade/pagemaker"
)

func main() {
	err := pagemaker.MakeWelcomePage("hyuki@example.com", "welcome.html")
	if err != nil {
		panic(err)
	}

}
