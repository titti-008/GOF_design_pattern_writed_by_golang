package pagemaker

import (
	"fmt"
	"os"
)

func MakeWelcomePage(email, filename string) (err error) {

	data, err := new(database).getData("./pagemaker/maildata.csv")
	if err != nil {
		return
	}

	username := data.getName(email)
	f, err := os.Create(filename)
	if err != nil {
		return
	}
	defer f.Close()
	w := newHtmlWriter(f)
	err = w.title(username + "'s web page")
	if err != nil {
		return
	}
	err = w.paragraph("Nice to meet you!")
	if err != nil {
		return
	}
	err = w.mailto(email, username)
	if err != nil {
		return
	}
	err = w.close()
	if err != nil {
		return
	}

	fmt.Println(filename, "is created for", email, " (", username, ")")
	return
}
