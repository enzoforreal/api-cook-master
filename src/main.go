package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"ApiCookMaster/src/handlers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to the Homepage!")
	})

	router.HandleFunc("/users", handlers.GetUsersHandler).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", handlers.GetuserHandler).Methods(http.MethodGet)
	router.HandleFunc("/users", handlers.CreateUserHandler).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", handlers.UpdateUserHandler).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", handlers.DeleteUserHandler).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8000", router))
}
