package models

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"memento/dto"
)

type User struct {
	ID        int8   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

type Users []User

func (u *User) GetUser(db *sql.DB) error {
	return db.QueryRow("SELECT username, first_name, last_name, email from User where id=$1", u.ID).Scan(&u.Username, &u.FirstName, &u.LastName, &u.Email)
}

func (u *User) UpdateUser(db *sql.DB) error {
	_, err := db.Exec("UPDATE User set username=$1,email=$2, first_name=$3, last_name=$4 where id=$5", u.Username, u.Email, u.FirstName, u.LastName, u.ID)

	return err
}

func (u *User) UpdatePassword(db *sql.DB) error {
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

func (u *User) DeleteUser(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (u *User) CreateUser(db *sql.DB) error {
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

func (u *User) LoginUser(db *sql.DB, loginUser dto.LoginUser) error {
	err := db.QueryRow("SELECT id, username, first_name, last_name, email, password  from User where username=$1", loginUser.Username).Scan(&u.ID, &u.Username, &u.FirstName, &u.LastName, &u.Email, &u.Password)

	if err != nil {
		return err
	}

	if !verifyPassword(loginUser.Password, u.Password) {
		return errors.New("username or password do not match")
	}

	return nil
}

func hashUserPassword(password string) []byte {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return hash
}

func verifyPassword(loginPassword string, hash string) bool {
	original_bytes := []byte(loginPassword)
	hash_bytes := []byte(hash)

	err := bcrypt.CompareHashAndPassword(hash_bytes, original_bytes)
	if err != nil {
		log.Fatal(err)
	}

	return true
}

func (u *User) ToDTO() dto.UserDTO {
	return dto.UserDTO{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
		Email:     u.Email,
	}
}
