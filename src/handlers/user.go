package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func getUsersFromDB() []User {

	return []User{
		{ID: 1, Name: "Grelet"},
		{ID: 2, Name: "Sananes"},
		{ID: 3, Name: "Bart"},
		{ID: 4, Name: "Dignat"},
	}
}

type User struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

var users []User

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := getUsersFromDB()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetuserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if r.Method != http.MethodGet {
		MethodNotAllowedHandler(w, r)
	}

	id, err := strconv.ParseInt(params["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	users = getUsersFromDB()

	found := false
	for _, user := range users {
		if user.ID == int32(id) {
			json.NewEncoder(w).Encode(user)
			found = true
			break
		}
	}

	if !found {
		w.WriteHeader(http.StatusNotFound)
		errorMessage := map[string]string{"error": "User not found"}
		json.NewEncoder(w).Encode(errorMessage)
	}
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	_ = json.NewDecoder(r.Body).Decode(&user)

	if r.Method != http.MethodPost {
		MethodNotAllowedHandler(w, r)
	}

	users = getUsersFromDB()

	highestID := int32(0)
	// l'ID utilisateur le plus grand
	for _, existingUser := range users {
		if existingUser.ID > highestID {
			highestID = existingUser.ID
		}
	}
	//Assigne au nouvel utilisateur le prochain ID disponible
	//pour garantir que chaque utilisateur auront un id unique
	user.ID = highestID + 1

	users = append(users, user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if r.Method != http.MethodPut {
		MethodNotAllowedHandler(w, r)
	}
	id, err := strconv.ParseInt(params["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	found := false
	for i, user := range users {
		if user.ID == int32(id) {
			users = append(users[:i], users[i+1:]...)

			var updateUser User
			err := json.NewDecoder(r.Body).Decode(&updateUser)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			defer r.Body.Close()

			updateUser.ID = int32(id)
			users = append(users, updateUser)
			json.NewEncoder(w).Encode(updateUser)
			found = true
			return
		}
	}

	if !found {
		w.WriteHeader(http.StatusNotFound)
		errorMessage := map[string]string{"error": "User not found"}
		json.NewEncoder(w).Encode(errorMessage)
	}
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if r.Method != http.MethodDelete {
		MethodNotAllowedHandler(w, r)
	}

	id, err := strconv.ParseInt(params["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	found := false
	for i, user := range users {
		if user.ID == int32(id) {
			users = append(users[:i], users[i+1:]...)
			found = true
			break
		}
	}

	if found {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	} else {
		w.WriteHeader(http.StatusNotFound)
		errorMessage := map[string]string{"error": "User not found"}
		json.NewEncoder(w).Encode(errorMessage)
	}
}
