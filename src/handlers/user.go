package handlers

import (
	"ApiCookMaster/src/auth"
	"ApiCookMaster/src/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
)

type User = models.User

var users []User

type UserHandler struct {
	userRepository UserRepository
	jwtManager     auth.JWTManager
}

func NewUserHandler(repository UserRepository) *UserHandler {
	return &UserHandler{userRepository: repository}
}

// methode SetJWTManager pour injecter le JWTManager dans UserHandler
func (j *UserHandler) SetJWTManager(jwtManager auth.JWTManager) {
	j.jwtManager = jwtManager
}

func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if r.Method != http.MethodPost {
		MethodNotAllowedHandler(w, r)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.userRepository.InsertUser(user.Nom, user.Prenom, user.Adresse, user.Email, user.Telephone, user.MotDepasse, user.PhotoDeProfil, user.EstAdmin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := h.userRepository.GetAllUsersFromDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetuserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if r.Method != http.MethodGet {
		MethodNotAllowedHandler(w, r)
	}

	id, err := strconv.ParseInt(params["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := h.userRepository.GetUserFromDB(int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		errorMessage := map[string]string{"error": "User not found"}
		json.NewEncoder(w).Encode(errorMessage)
	} else {
		json.NewEncoder(w).Encode(user)
	}
}

func (h *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if r.Method != http.MethodPut {
		MethodNotAllowedHandler(w, r)
		return
	}

	id, err := strconv.ParseInt(params["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	user.ID = int32(id)
	updatedUser, err := h.userRepository.UpdateUserFromDB(user.ID, user.Nom, user.Prenom, user.Adresse, user.Email, user.Telephone, user.MotDepasse, user.PhotoDeProfil, user.EstAdmin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if updatedUser == nil {
		w.WriteHeader(http.StatusNotFound)
		errorMessage := map[string]string{"error": "User not found"}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedUser)
}

func (h *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if r.Method != http.MethodDelete {
		MethodNotAllowedHandler(w, r)
	}

	id, err := strconv.ParseInt(params["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.userRepository.DeleteUserFromDB(int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		MethodNotAllowedHandler(w, r)
		return
	}

	var credentials struct {
		Email      string `json:"email"`
		MotDepasse string `json:"mot_de_passe"`
	}

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	log.Printf("Received login request for email: %s\n", credentials.Email)

	user, err := h.userRepository.GetUserByEmail(credentials.Email)
	if err != nil {
		// Handle error here
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user == nil || bcrypt.CompareHashAndPassword([]byte(user.MotDepasse), []byte(credentials.MotDepasse)) != nil {
		log.Printf("Invalid email or password for email: %s\n", credentials.Email)
		w.WriteHeader(http.StatusUnauthorized)
		errorMessage := map[string]string{"error": "Invalid email or password"}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	log.Printf("Successful login for email: %s\n", credentials.Email)

	token, err := h.jwtManager.Generate(*user)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
