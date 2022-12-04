package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("hello flyweight")

	flag.Parse()
	arg := flag.Arg(0)

	if arg == "" {
		panic(fmt.Errorf("Please input digits. example: '1212123'. but: args is: %v", arg))
	}

	bs, err := newBigString(arg)
	if err != nil {
		panic(err)
	}

	bs.print()
}
