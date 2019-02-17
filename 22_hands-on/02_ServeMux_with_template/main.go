package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func root(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, "This is the data")
	//err := tpl.Execute(w, req)
	if err != nil {
		log.Fatalln(err)
	}	++







}

func dog(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, "This is the data for dog")
	if err != nil {
		log.Fatalln(err)
	}
}

func me(w http.ResponseWriter, req *http.Request) {
	//err := tpl.Execute(w, "This is the data for me")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	io.WriteString(w,"David Marubbi")
}

var tpl *template.Template

func init() {
	 tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {


	http.HandleFunc("/", root)

	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}
