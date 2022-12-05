package main

import (
	"fmt"
	"time"
)

type printer struct {
	name string
}

var _ printable = new(printer)

func newPrinter(name string) *printer {
	p := new(printer)
	p.name = name
	p.heavyJob("Printerのインスタンスを生成中")

	return p
}

func (p *printer) setPrinterName(name string) {
	p.name = name
}

func (p printer) getPrinterName() string {
	return p.name
}

func (p printer) print(str string) {
	fmt.Println("=== " + p.name + " ===")
	fmt.Println(str)
}

func (p printer) heavyJob(msg string) {
	fmt.Println(msg)
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println(".")
	}
	fmt.Println("完了。")
}
