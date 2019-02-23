package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct{ Name, Address, City, Zip, Region string }

type hotels []hotel

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("dog.gohtml"))
}

func main() {
	h := hotels{
		hotel{
			Name:    "The Ramarda",
			Address: "12 Lake Drive",
			City:    "Denver",
			Zip:     "AB120",
			Region:  "Central"},
		hotel{
			"The Ritz",
			"1 Main Street",
			"New York",
			"NY100",
			"Northern"},
	}

	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}
}
