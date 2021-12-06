package main

import (
	"os"
	"text/template"
)

const (
	templateName   = "test"
	templateFormat = "Součet {{.X}} + {{.Y}} = {{.Z}}"
)

type User struct {
	FirstName string
	Surname   string
	Born      string
}

func main() {
	// vytvoření nové šablony
	tmpl := template.Must(template.New(templateName).Parse(templateFormat))

	// tyto hodnoty budou použity při aplikaci šablony
	user := User{
		FirstName: "Jára",
		Surname:   "Cimrman",
		Born:      "Böhmen",
	}

	// aplikace šablony - přepis hodnot
	err := tmpl.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
