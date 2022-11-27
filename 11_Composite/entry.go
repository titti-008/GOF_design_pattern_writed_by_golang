package main

type entryInterface interface {
    getName() string
    getSize() int
    printList(prefix string)
}

