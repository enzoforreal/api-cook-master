package main

import (
	"ApiCookMaster/src/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to the Homepage!")
	})

	userRepository := handlers.NewUserRepository()         // créer une instance UserRepository
	userHandler := handlers.NewUserHandler(userRepository) // utilise UserRepository dans UserHandler
	router.HandleFunc("/users", userHandler.GetUsersHandler).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", userHandler.GetuserHandler).Methods(http.MethodGet)
	router.HandleFunc("/users", userHandler.CreateUserHandler).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", userHandler.UpdateUserHandler).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", userHandler.DeleteUserHandler).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8000", router))
}
