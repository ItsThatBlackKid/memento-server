package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"memento/controller"
	"memento/utils"
	"net/http"
	"os"
	"strings"
)

func handleAuthError(w http.ResponseWriter, reason string) {
	controller.RespondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Unauthorized Access: %s", reason))
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")
			tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
			if tokenString == "" {
				controller.RespondWithError(
					w,
					http.StatusUnauthorized,
					"Unauthorised access: missing bearer token",
				)
				return
			}

			token, err := jwt.ParseWithClaims(
				tokenString,
				&utils.UserClaims{},
				func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, errors.New("unexpected signing method")
					}

					return []byte(os.Getenv("TOKEN_SECRET")), nil
				},
			)

			if err != nil || !token.Valid {
				log.Println("Invalid token error", err.Error())
				handleAuthError(w, "invalid token")
				return
			}

			claims, ok := token.Claims.(*utils.UserClaims)
			if !ok {
				handleAuthError(w, "invalid claims")
				return
			}

			ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
			next.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}
