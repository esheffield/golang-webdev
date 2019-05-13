package main

import (
	"io"
	"net/http"
)

func dogHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "doggy doggy doggy dog")
}

func catHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "kitty kitty kitty cat")
}

func main() {

	http.HandleFunc("/dog/", dogHandler)
	http.HandleFunc("/cat", catHandler)

	http.ListenAndServe(":8080", nil)
}
