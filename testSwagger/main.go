package main

import (
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	_ "testProj/testSwagger/docs"
)

type UserHandler struct {
}

// @Summary Create
// @Description create user
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{} "User created successfully"
// @Router /users [post]
func (u *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create user "))
}
func Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create user "))
}

// @title API Title
// @version 1.0
// @description This is a sample server.

// @host localhost:8080
// @BasePath /
// @schemes http

func main() {

	r := chi.NewRouter()
	r.Get("/user", Create)
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	http.ListenAndServe(":8080", r)
}
