package main

import (
	"ApiCookMaster/src/handlers"
	"ApiCookMaster/src/handlers/db"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// Appel la fonction init() du package db pour charger les variables d'environnement
	db.Init()
	router := mux.NewRouter()
	router.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome !")
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
