package handlers

import (
	"ApiCookMaster/src/handlers/db"
	"time"
)

type UserRepository interface {
	InsertUser(nom, prenom, adresse, email, telephone, mot_de_passe, photo_de_profil string, est_admin bool) error
	GetUserFromDB(id int32) (*User, error)
	GetAllUsersFromDB() ([]User, error)
	UpdateUserFromDB(id int32, nom string, prenom string, adresse string, email string, telephone string, mot_de_passe string, photo_de_profil string, est_admin bool) (*User, error)
	DeleteUserFromDB(id int32) error
}

type ReservationRepository interface {
	InsertReservation(IDUser int32, IDEvenement int32, IDCours int32, DateReservation time.Time, NbPersonnes int32, StatusReservation string) error
	GetReservationFromDB(id int32) (*Reservation, error)
	GetAllReservationsFromDB() ([]Reservation, error)
	UpdateReservationFromDB(id int32, IDUser int32, IDEvenement int32, IDCours int32, DateReservation time.Time, NbPersonnes int32, StatusReservation string) (*Reservation, error)
	DeleteReservationFromDB(id int32) error
}

type reservationRepositoryImpl struct{}

func NewReservationRepository() *reservationRepositoryImpl {
	return &reservationRepositoryImpl{}
}

type userRepositoryImpl struct{}

func NewUserRepository() *userRepositoryImpl {
	return &userRepositoryImpl{}
}

func (*userRepositoryImpl) InsertUser(nom, prenom, adresse, email, telephone, mot_de_passe, photo_de_profil string, est_admin bool) error {
	return db.InsertUser(nom, prenom, adresse, email, telephone, mot_de_passe, photo_de_profil, est_admin)
}

func (*userRepositoryImpl) GetAllUsersFromDB() ([]User, error) {
	return db.GetAllUsersFromDB()
}

func (*userRepositoryImpl) GetUserFromDB(id int32) (*User, error) {
	return db.GetUserFromDB(id)
}

func (*userRepositoryImpl) UpdateUserFromDB(id int32, nom string, prenom string, adresse string, email string, telephone string, mot_de_passe string, photo_de_profil string, est_admin bool) (*User, error) {
	return db.UpdateUserFromDB(id, nom, prenom, adresse, email, telephone, mot_de_passe, photo_de_profil, est_admin)
}

func (*userRepositoryImpl) DeleteUserFromDB(id int32) error {
	return db.DeleteUserFromDB(id)
}

func (*reservationRepositoryImpl) InsertReservation(IDUser int32, IDEvenement int32, IDCours int32, DateReservation time.Time, NbPersonnes int32, StatusReservation string) error {
	return db.InsertReservation(IDUser, IDEvenement, IDCours, DateReservation, NbPersonnes, StatusReservation)
}

func (*reservationRepositoryImpl) GetAllReservationsFromDB() ([]Reservation, error) {
	return db.GetAllReservationsFromDB()
}

func (*reservationRepositoryImpl) GetReservationFromDB(id int32) (*Reservation, error) {
	return db.GetReservationFromDB(id)
}

func (*reservationRepositoryImpl) UpdateReservationFromDB(id int32, IDUser int32, IDEvenement int32, IDCours int32, DateReservation time.Time, NbPersonnes int32, StatusReservation string) (*Reservation, error) {
	return db.UpdateReservationFromDB(id, IDUser, IDEvenement, IDCours, DateReservation, NbPersonnes, StatusReservation)
}

func (*reservationRepositoryImpl) DeleteReservationFromDB(id int32) error {
	return db.DeleteReservationFromDB(id)
}
