package main

import (
	"fmt"
)

func main() {
	p := newPrinterProxy("alice")
	fmt.Println("名前は現在" + p.getPrinterName() + "です。")
	p.setPrinterName("Bob")
	fmt.Println("名前は現在" + p.getPrinterName() + "です。")
	p.print("Hello, world.")
}
