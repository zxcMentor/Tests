package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Get("/maths", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("ser"))
	})
	http.ListenAndServe(":8080", r)
}
