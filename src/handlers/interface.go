package handlers

import (
	"ApiCookMaster/src/handlers/db"
)

type UserRepository interface {
	InsertUser(nom, prenom, adresse, email, telephone, mot_de_passe, photo_de_profil string, est_admin bool) error
	GetUserFromDB(id int32) (*User, error)
	GetAllUsersFromDB() ([]User, error)
	UpdateUserFromDB(id int32, nom string, prenom string, adresse string, email string, telephone string, mot_de_passe string, photo_de_profil string, est_admin bool) (*User, error)
	DeleteUserFromDB(id int32) error
	GetUserByEmail(email string) (*User, error)
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

func (*userRepositoryImpl) GetUserByEmail(email string) (*User, error) {
	return db.GetUserByEmail(email)
}
