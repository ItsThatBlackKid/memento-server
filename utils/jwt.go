package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"memento/dto"
)

type UserJwtClaim struct {
	jwt.Claims
	User *dto.UserDTO `json:"user"`
}

func EncodeJwt(user dto.UserDTO) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserJwtClaim{
		Claims: nil,
		User:   &user,
	})

	return token.SignedString([]byte("shhh"))
}

func DecodeJwt(tokenString string) (*dto.UserDTO, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("shh"), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(UserJwtClaim)
	if !ok {
		return nil, errors.New("invalid JWT Token")
	}

	return claims.User, err
}

func VerifyJwt(tokenString string) error {
	_, err := DecodeJwt(tokenString)
	return err
}
