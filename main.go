package main

import (
	"net/http"

	"github.com/bjoernschilberg/special-system/handler"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/users", handler.GetUsers)
	mux.HandleFunc("/users/{id}", handler.UserByID)
	mux.HandleFunc("POST /users", handler.CreateUser)
	mux.HandleFunc("DELETE /users/{id}", handler.DeleteUser)

	http.ListenAndServe(":8888", mux)

}
