package main

import (
	"ApiCookMaster/src/handlers"
	"ApiCookMaster/src/handlers/db"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

	reservationRepository := handlers.NewReservationRepository()                // créer une instance UserRepository
	reservationHandler := handlers.NewReservationHandler(reservationRepository) // utilise UserRepository dans UserHandler
	router.HandleFunc("/reservations", reservationHandler.GetReservationsHandler).Methods(http.MethodGet)
	router.HandleFunc("/reservations/{id}", reservationHandler.GetreservationHandler).Methods(http.MethodGet)
	router.HandleFunc("/reservations", reservationHandler.CreateReservationHandler).Methods(http.MethodPost)
	router.HandleFunc("/reservations/{id}", reservationHandler.UpdateReservationHandler).Methods(http.MethodPut)
	router.HandleFunc("/reservations/{id}", reservationHandler.DeleteReservationHandler).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8000", router))
}
