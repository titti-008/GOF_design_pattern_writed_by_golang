package main

type printable interface {
	setPrinterName(name string)
	getPrinterName() string
	print(str string)
}
