package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"memento/dto"
	"os"
	"time"
)

type UserClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func CreateJwtForUser(user dto.UserDTO) (string, error) {
	claims := UserClaims{
		user.ID,
		user.Username,
		user.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "memento-server",
			Subject:   user.Username,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(os.Getenv("TOKEN_SECRET"))
}
