package main

import (
	"library/src/main/book"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /books", book.FindAll)
	mux.HandleFunc("GET /books/{id}", book.FindOne)
	mux.HandleFunc("POST /books", book.Create)
	mux.HandleFunc("PUT /books/{id}", book.Update)
	mux.HandleFunc("DELETE /books/{id}", book.Delete)

	if err := http.ListenAndServe(":80", mux); err != nil {
		panic(err)
	}
}
