package models

import "time"

type Reservation struct {
	ID                int32     `json:"id"`
	IDUser            int32     `json:"id_user"`
	IDEvenement       int32     `json:"id_evenement"`
	IDCours           int32     `json:"id_cours"`
	DateReservation   time.Time `json:"date_reservation"`
	NbPersonnes       int32     `json:"nb_personnes"`
	StatusReservation string    `json:"status_reservation"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
