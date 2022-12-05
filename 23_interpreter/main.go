package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	length, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	w := newWindow(length, position{X: length / 2, Y: length / 2})
	for scanner.Scan() {
		fmt.Println("Input: ")
		line := scanner.Text()
		if line == "" {
			break
		}
		strArr := strings.Split(line, " ")

		for i, str := range strArr {
			fmt.Printf("%v: %v\n", i, str)
		}

		context := newContext(line, w)
		node := new(programNode)

		err := node.parse(context)
		if err != nil {
			panic(err)
		}
		fmt.Println(node)

		w.print()
		fmt.Print("input ->: ")
	}
}
