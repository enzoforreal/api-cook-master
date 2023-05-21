package db

import (
	"ApiCookMaster/src/models"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Reservation = models.Reservation

func InsertReservation(IDUser int32, IDEvenement int32, IDCours int32, DateReservation time.Time, NbPersonnes int32, StatusReservation string) error {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return fmt.Errorf("Error opening database connection: %w", err)
	}
	defer db.Close()

	stmt, err := db.Prepare(`
        INSERT INTO reservations (id_user, id_evenement, id_cours, date_reservation, nb_personnes, status_reservation , created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, now(), now())
        RETURNING id, id_user, id_evenement, id_cours, date_reservation, nb_personnes, status_reservation, created_at, updated_at
    `)
	if err != nil {
		return fmt.Errorf("Error preparing sql statement: %w", err)
	}

	var reservation Reservation
	err = stmt.QueryRow(IDUser, IDEvenement, IDCours, DateReservation, NbPersonnes, StatusReservation).Scan(&reservation.ID, &reservation.IDUser, &reservation.IDEvenement, &reservation.IDCours, &reservation.DateReservation, &reservation.NbPersonnes, &reservation.StatusReservation, &reservation.CreatedAt, &reservation.UpdatedAt)
	if err != nil {
		return fmt.Errorf("Error executing sql statement: %w", err)
	}

	return nil
}

func GetAllReservationsFromDB() ([]Reservation, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("Error opening database connection: %w", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, id_user, id_evenement, id_cours, date_reservation, nb_personnes, status_reservation FROM reservations")
	if err != nil {
		return nil, fmt.Errorf("Error querying database: %w", err)
	}
	defer rows.Close()

	var reservations []Reservation
	for rows.Next() {
		var reservation Reservation
		err := rows.Scan(&reservation.IDUser, &reservation.IDEvenement, &reservation.IDCours, &reservation.DateReservation, &reservation.NbPersonnes, &reservation.StatusReservation)
		if err != nil {
			return nil, fmt.Errorf("Error scanning row: %w", err)
		}
		reservations = append(reservations, reservation)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterating over rows: %w", err)
	}

	return reservations, nil
}

func GetReservationFromDB(id int32) (*Reservation, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("Error opening database connection: %w", err)
	}
	defer db.Close()

	var reservation Reservation
	err = db.QueryRow("SELECT id, id_user, id_evenement, id_cours, date_reservation, nb_personnes, status_reservation FROM reservations WHERE id = $1", id).Scan(&reservation.ID, &reservation.IDUser, &reservation.IDEvenement, &reservation.IDCours, &reservation.DateReservation, &reservation.NbPersonnes, &reservation.StatusReservation)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("Error retrieving reservation from database: %w", err)
	}

	return &reservation, nil
}

func UpdateReservationFromDB(ID int32, IDUser int32, IDEvenement int32, IDCours int32, DateReservation time.Time, NbPersonnes int32, StatusReservation string) (*Reservation, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("Error opening database connection: %w", err)
	}
	defer db.Close()

	var reservation Reservation
	err = db.QueryRow("UPDATE reservations SET id_user=$1, id_evenement=$2, id_cours=$3, date_reservation=$4, nb_personnes=$5, status_reservation=$6 WHERE id=$7 RETURNING id, id_user, id_evenement, id_cours, date_reservation, nb_personnes, status_reservation",
		IDUser, IDEvenement, IDCours, DateReservation, NbPersonnes, StatusReservation, ID).Scan(&reservation.IDUser, &reservation.IDEvenement, &reservation.IDCours, &reservation.DateReservation, &reservation.NbPersonnes, &reservation.StatusReservation, &reservation.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("Error updating reservation from database: %w", err)
	}

	return &reservation, nil
}

func DeleteReservationFromDB(id int32) error {
	db, err := sql.Open("postgres", dataSourceName)

	if err != nil {
		return fmt.Errorf("Error opening database connection: %w", err)
	}
	defer db.Close()

	// Vérifier que la reservation existe avant de la supprimer
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM reservations WHERE id = $1", id).Scan(&count)
	if err != nil {
		return fmt.Errorf("Error querying database: %w", err)
	}
	if count == 0 {
		return fmt.Errorf("Reservation not found")
	}

	_, err = db.Exec("DELETE FROM reservations WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("Error executing sql statement : %w", err)
	}

	return nil
}
