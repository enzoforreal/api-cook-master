package main

import (
	"ApiCookMaster/src/auth"
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

	// instance de JWTManager
	jwtManager := auth.NewJWT()

	router := mux.NewRouter()
	router.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome !")
	})

	userRepository := handlers.NewUserRepository() // créer une instance UserRepository
	recipeRepository := db.NewRecipeRepository()
	userHandler := handlers.NewUserHandler(userRepository, recipeRepository) // utilise UserRepository dans UserHandler
	userHandler.SetJWTManager(jwtManager)                                    // Set JWT Manager to UserHandler

	// Create subrouter for routes requiring JWT authentication
	authRouter := router.PathPrefix("/").Subrouter()
	authRouter.Use(auth.MiddlewareJWT(jwtManager))
	// Add login route
	authRouter.HandleFunc("/login", userHandler.LoginHandler).Methods(http.MethodPost)
	// Move your existing routes to the subrouter
	authRouter.HandleFunc("/users", userHandler.GetUsersHandler).Methods(http.MethodGet)
	authRouter.HandleFunc("/users/{id}", userHandler.GetuserHandler).Methods(http.MethodGet)
	authRouter.HandleFunc("/users", userHandler.CreateUserHandler).Methods(http.MethodPost)
	authRouter.HandleFunc("/users/{id}", userHandler.UpdateUserHandler).Methods(http.MethodPut)
	authRouter.HandleFunc("/users/{id}", userHandler.DeleteUserHandler).Methods(http.MethodDelete)

	//endpoint pour recuperer les recettes aleatoires
	authRouter.HandleFunc("/recipes/random", userHandler.GetRecipeHandler).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8000", router))
}
