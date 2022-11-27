package main

import (
    "fmt"
)

func main() {
    fmt.Println("Making root entries...")
    rootDir := newDir("root")
    bin := newDir("bin")
    tmp := newDir("tmp")
    usr := newDir("usr")

    rootDir.add(bin)
    rootDir.add(tmp)
    rootDir.add(usr)

    bin.add(newFile("vi", 10000))
    bin.add(newFile("latex", 20000))
    rootDir.printList("~")
    fmt.Println()


    fmt.Println("Making user entries...")
    yuki := newDir("yuki")
    hanako := newDir("hanako")
    tomura := newDir("tomura")

    usr.add(yuki)
    usr.add(hanako)
    usr.add(tomura)
    yuki.add(newFile("diary.html", 100))
    yuki.add(newFile("Composite.go", 200))
    hanako.add(newFile("memo.txt", 300))
    tomura.add(newFile("game.doc", 400))
    tomura.add(newFile("junk.mail", 500))

    rootDir.printList("~")



}
