package main

import (
	"io"
	"net/http"
)

func root(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "This is the root")
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "This is the dog page")
}

func me(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hi, my name is Dave")
}

func main() {

	http.Handle("/", http.HandlerFunc(root))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}
