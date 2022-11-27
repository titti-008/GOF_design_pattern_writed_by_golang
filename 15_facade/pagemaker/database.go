package pagemaker

import (
	"encoding/csv"
	"os"
)

type database struct {
}

type parson struct {
	name  string
	email string
}

type data []parson

func (d *database) getData(dbname string) (data, error) {

	f, err := os.Open(dbname)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	rows, err := r.ReadAll()

	result := data{}

	for _, row := range rows[1:] {
		p := new(parson)
		p.email = row[0]
		p.name = row[1]
		result = append(result, *p)
	}

	return result, nil
}

func (d data) getName(email string) string {
	for _, p := range d {
		if p.email == email {
			return p.name
		}
	}
	return ""
}
