package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("public"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("./templates/index.gohtml"))
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
