package models

import "time"

type User struct {
	ID            int32     `json:"id"`
	Nom           string    `json:"nom"`
	Prenom        string    `json:"prenom"`
	Adresse       string    `json:"adresse"`
	Email         string    `json:"email"`
	Telephone     string    `json:"telephone"`
	MotDepasse    string    `json:"mot_de_passe"`
	PhotoDeProfil string    `json:"photoDeProfil"`
	EstAdmin      bool      `json:"est_admin"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
