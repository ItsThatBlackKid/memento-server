package models

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	ID        int16  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

type Users []User

func hashUserPassword(password string) []byte {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return hash
}

func (u *User) getUser(db *sql.DB) error {
	return db.QueryRow("SELECT username, first_name, last_name, email from User where id=$1", u.ID).Scan(&u.Username, &u.FirstName, &u.LastName, &u.Email)
}

func (u *User) updateUser(db *sql.DB) error {
	_, err := db.Exec("UPDATE User set username=$1,email=$2, first_name=$3, last_name=$4 where id=$5", u.Username, u.Email, u.FirstName, u.LastName, u.ID)

	return err
}

func (u *User) updatePassword(db *sql.DB) error {
	if u.Password == "" {
		return errors.New("Password is empty")
	}

	u.Password = string(hashUserPassword(u.Password))

	_, err := db.Exec("UPDATE User set password=$1 where id=$1", u.Password, u.ID)

	if err != nil {
		return err
	}

	return nil
}

func (u *User) deleteUser(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (u *User) createUser(db *sql.DB) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	u.Password = string(hash)

	if err != nil {
		return err
	}

	err = db.QueryRow(
		"INSERT INTO User(username, email, first_name, last_name, password) VALUES ($1,$2,$3,$4,$5) RETURNING id",
		u.Username, u.Email, u.FirstName, u.LastName, u.Password).Scan(&u.ID)

	if err != nil {
		return err
	}

	return nil
}
