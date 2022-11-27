package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello visitor")

	fmt.Println("====== makeing root entries ======")
	rootDir := newDir("root")
	bin := newDir("bin")
	tmp := newDir("tmp")
	usr := newDir("usr")
	rootDir.add(bin)
	rootDir.add(tmp)
	rootDir.add(usr)
	bin.add(newFile("vi", 10000))
	bin.add(newFile("latex", 20000))
	rootDir.accept(newListVisitor())
	fmt.Println()

	fmt.Println("====== makeing root entries ======")
	yuki := newDir("yuki")
	hanako := newDir("hanako")
	tomura := newDir("tomura")
	usr.add(yuki)
	usr.add(hanako)
	usr.add(tomura)
	yuki.add(newFile("diary.html", 100))
	yuki.add(newFile("Composite.go", 200))
	hanako.add(newFile("memo.tex", 300))
	tomura.add(newFile("game.html", 400))
	tomura.add(newFile("junk.mail", 500))

	ffv := newFFV(".html")
	rootDir.accept(ffv)

	fmt.Println("HTML files are:")
	for _, f := range ffv.getFounds() {
		fmt.Println(f)
	}

	rootDir.accept(new(sizeVisitor))
	rootDir.accept(newListVisitor())
}
