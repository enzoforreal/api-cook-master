package db

import (
	"ApiCookMaster/src/models"
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type User = models.User

const dataSourceName = "user=postgres password=secret host=localhost dbname=cookmaster sslmode=disable"

func InsertUser(nom, prenom, adresse, email, telephone, mot_de_passe, photo_de_profil string, est_admin bool) error {
	db, err := sql.Open("postgres", "dataSourceName")
	if err != nil {
		return fmt.Errorf("Error opening database connection: %w", err)
	}
	defer db.Close()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(mot_de_passe), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("Error hashing password: %w", err)
	}

	stmt, err := db.Prepare(`
        INSERT INTO users (nom, prenom, adresse, email, telephone, mot_de_passe, photo_de_profil, est_admin, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, now(), now())
        RETURNING id, created_at, updated_at
    `)
	if err != nil {
		return fmt.Errorf("Error preparing sql statement: %w", err)
	}

	var user User
	err = stmt.QueryRow(nom, prenom, adresse, email, telephone, hashedPassword, photo_de_profil, est_admin).Scan(
		&user.ID, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("Error executing sql statement: %w", err)
	}

	return nil
}

func GetAllUsersFromDB() ([]User, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("Error opening database connection: %w", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, nom, prenom, adresse, email, telephone, mot_de_passe, photo_de_profil, est_admin FROM users")
	if err != nil {
		return nil, fmt.Errorf("Error querying database: %w", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Nom, &user.Prenom, &user.Adresse, &user.Email, &user.Telephone, &user.MotDepasse, &user.PhotoDeProfil, &user.EstAdmin)
		if err != nil {
			return nil, fmt.Errorf("Error scanning row: %w", err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterating over rows: %w", err)
	}

	return users, nil
}

func GetUserFromDB(id int32) (*User, error) {
	db, err := sql.Open("postgres", "dataSourceName")
	if err != nil {
		return nil, fmt.Errorf("Error opening database connection: %w", err)
	}
	defer db.Close()

	var user User
	err = db.QueryRow("SELECT id, nom, prenom, adresse, email, telephone, mot_de_passe, photo_de_profil, est_admin FROM users WHERE id = $1", id).Scan(&user.ID, &user.Nom, &user.Prenom, &user.Adresse, &user.Email, &user.Telephone, &user.MotDepasse, &user.PhotoDeProfil, &user.EstAdmin)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("Error retrieving user from database: %w", err)
	}

	return &user, nil
}

func UpdateUserFromDB(id int32, nom string, prenom string, adresse string, email string, telephone string, mot_de_passe string, photo_de_profil string, est_admin bool) (*User, error) {
	db, err := sql.Open("postgres", "dataSourceName")
	if err != nil {
		return nil, fmt.Errorf("Error opening database connection: %w", err)
	}
	defer db.Close()

	var user User
	err = db.QueryRow("UPDATE users SET nom=$1, prenom=$2, adresse=$3, email=$4, telephone=$5, mot_de_passe=$6, photo_de_profil=$7, est_admin=$8 WHERE id=$9 RETURNING id, nom, prenom, adresse, email, telephone, mot_de_passe, photo_de_profil, est_admin",
		nom, prenom, adresse, email, telephone, mot_de_passe, photo_de_profil, est_admin, id).Scan(&user.ID, &user.Nom, &user.Prenom, &user.Adresse, &user.Email, &user.Telephone, &user.MotDepasse, &user.PhotoDeProfil, &user.EstAdmin)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		if strings.Contains(err.Error(), "users_email_key") {
			return nil, fmt.Errorf("Email already exists")
		}
		return nil, fmt.Errorf("Error updating user from database: %w", err)
	}

	return &user, nil
}

func DeleteUserFromDB(id int32) error {
	db, err := sql.Open("postgres", "dataSourceName")
	if err != nil {
		return fmt.Errorf("Error opening database connection: %w", err)
	}
	defer db.Close()

	// Vérifier que l'utilisateur existe avant de le supprimer
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE id = $1", id).Scan(&count)
	if err != nil {
		return fmt.Errorf("Error querying database: %w", err)
	}
	if count == 0 {
		return fmt.Errorf("User not found")
	}

	_, err = db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("Error executing sql statement : %w", err)
	}

	return nil
}
