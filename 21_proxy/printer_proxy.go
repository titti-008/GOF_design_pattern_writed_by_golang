package main

type printerProxy struct {
	name string
	real printable
}

var _ printable = new(printerProxy)

func newPrinterProxy(name string) *printerProxy {
	pp := new(printerProxy)
	pp.name = name
	pp.real = nil
	return pp
}

func (pp *printerProxy) setPrinterName(name string) {
	if pp.real != nil {
		pp.real.setPrinterName(name)
	}
	pp.name = name
}

func (pp *printerProxy) getPrinterName() string {
	return pp.name
}

func (pp *printerProxy) print(str string) {
	pp.realize()
	pp.real.print(str)
}

func (pp *printerProxy) realize() {
	if pp.real == nil {
		pp.real = newPrinter(pp.name)
	}
}
