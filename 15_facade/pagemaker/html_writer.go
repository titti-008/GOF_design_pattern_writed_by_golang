package pagemaker

import (
	"fmt"
	"io"
)

type htmlWriter struct {
	writer io.Writer
}

func newHtmlWriter(w io.Writer) *htmlWriter {
	hw := new(htmlWriter)
	hw.writer = w
	return hw
}

func (w *htmlWriter) title(title string) error {
	str := ""
	str += "<!DOCTYPE html>\n"
	str += "<html>\n"
	str += "<head>\n"
	str += fmt.Sprintf("<title>%v</title>\n", title)
	str += "</head>\n"
	str += "<body>\n"
	str += fmt.Sprintf("<h1>%v</h1>\n", title)
	_, err := w.writer.Write([]byte(str))
	return err
}

func (w *htmlWriter) paragraph(msg string) error {
	str := fmt.Sprintf("<p>%v</p>\n", msg)
	_, err := w.writer.Write([]byte(str))
	return err
}

func (w *htmlWriter) link(href, caption string) error {
	l := fmt.Sprintf("<a href=\"%v\">%v</a>", href, caption)
	err := w.paragraph(l)
	return err
}

func (w *htmlWriter) mailto(email, name string) error {
	href := fmt.Sprintf("mailto: %v", email)
	err := w.link(href, name)
	return err
}

func (w *htmlWriter) close() error {
	str := ""
	str += "</body>\n"
	str += "</html>\n"
	_, err := w.writer.Write([]byte(str))
	return err
}
