package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct { Name, Address, City, Zip, Region string }

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotels := []hotel{
		{"The Ramarda", "12 Lake Drive", "Denver", "AB120", "Central"},
		{"The Ritz", "1 Main Street", "New York", "NY100", "Northern"},
	}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
