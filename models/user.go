package models

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"memento/context"
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

func (u *User) GetUser() error {
	if result := context.Context.DB.First(&u, u.ID); result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *User) UpdateUser() error {
	result := context.Context.DB.Model(&u).Select(
		"first_name",
		"last_name",
		"email",
		"username",
	).Updates(&u)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *User) UpdatePassword() error {
	if u.Password == "" {
		return errors.New("password is empty")
	}

	u.Password = string(hashUserPassword(u.Password))

	result := context.Context.DB.Model(&u).Update("password", u.Password)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *User) DeleteUser() error {
	return errors.New("not implemented")
}

func (u *User) CreateUser() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	u.Password = string(hash)

	if err != nil {
		return err
	}

	result := context.Context.DB.Create(&u)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *User) LoginUser(loginUser dto.LoginUser) error {
	result := context.Context.DB.First(&u, "username=$1", loginUser.Username)

	if result.Error != nil {
		return result.Error
	}

	if !verifyPassword(loginUser.Password, u.Password) {
		return errors.New("username or password do not match")
	}

	return nil
}

func hashUserPassword(password string) []byte {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		log.Fatal(err)
	}
	return hash
}

func verifyPassword(loginPassword string, hash string) bool {
	originalBytes := []byte(loginPassword)
	hashBytes := []byte(hash)

	err := bcrypt.CompareHashAndPassword(hashBytes, originalBytes)
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
