package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Decrator")

	b1 := newStringDisplay("Hello, world")
	b2 := newSideBorder(b1, "#")
	b3 := newFullBorder(b2)

	show(b1)
	show(b2)
	show(b3)
	b4 := newSideBorder(
		newFullBorder(
			newFullBorder(
				newSideBorder(
					newFullBorder(
						b1,
					),
					"*",
				),
			),
		),
		"/",
	)
	show(b4)

	md := newMultiStringDisplay()
	md.add("line1")
	md.add("line2 t e s t")
	md.add("line3")
	b5 := newSideBorder(
		newFullBorder(
			newFullBorder(
				newSideBorder(
					newFullBorder(
						md,
					),
					"*",
				),
			),
		),
		"/",
	)

	show(b5)
}
