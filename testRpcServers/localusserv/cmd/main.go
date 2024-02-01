package main

import (
	"github.com/go-chi/chi"
	"net/http"
	"testProj/testRpcServers/localusserv/controller"
)

func main() {

	r := chi.NewRouter()
	mh := controller.NewMathHand()

	r.Get("/math", mh.PlusHandler)

	http.ListenAndServe(":8080", r)
}
