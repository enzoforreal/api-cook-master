package handlers

import (
	"ApiCookMaster/src/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Reservation = models.Reservation

type ReservationHandler struct {
	reservationRepository ReservationRepository
}

func NewReservationHandler(repository ReservationRepository) *ReservationHandler {
	return &ReservationHandler{reservationRepository: repository}
}

func (h *ReservationHandler) CreateReservationHandler(w http.ResponseWriter, r *http.Request) {
	var reservation Reservation
	if r.Method != http.MethodPost {
		MethodNotAllowedHandler(w, r)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&reservation)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.reservationRepository.InsertReservation(reservation.IDUser, reservation.IDEvenement, reservation.IDCours, reservation.DateReservation, reservation.NbPersonnes, reservation.StatusReservation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(reservation)
}

func (h *ReservationHandler) GetReservationsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		MethodNotAllowedHandler(w, r)
	}

	reservations, err := h.reservationRepository.GetAllReservationsFromDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reservations)
}

func (h *ReservationHandler) GetreservationHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if r.Method != http.MethodGet {
		MethodNotAllowedHandler(w, r)
	}

	id, err := strconv.ParseInt(params["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	reservation, err := h.reservationRepository.GetReservationFromDB(int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if reservation == nil {
		w.WriteHeader(http.StatusNotFound)
		errorMessage := map[string]string{"error": "Reservation not found"}
		json.NewEncoder(w).Encode(errorMessage)
	} else {
		json.NewEncoder(w).Encode(reservation)
	}
}

func (h *ReservationHandler) UpdateReservationHandler(w http.ResponseWriter, r *http.Request) {
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

	var reservation Reservation
	err = json.NewDecoder(r.Body).Decode(&reservation)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	reservation.ID = int32(id)
	updatedReservation, err := h.reservationRepository.UpdateReservationFromDB(reservation.ID, reservation.IDUser, reservation.IDEvenement, reservation.IDCours, reservation.DateReservation, reservation.NbPersonnes, reservation.StatusReservation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if updatedReservation == nil {
		w.WriteHeader(http.StatusNotFound)
		errorMessage := map[string]string{"error": "Reservation not found"}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedReservation)
}

func (h *ReservationHandler) DeleteReservationHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if r.Method != http.MethodDelete {
		MethodNotAllowedHandler(w, r)
	}

	id, err := strconv.ParseInt(params["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.reservationRepository.DeleteReservationFromDB(int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
