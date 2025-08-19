package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/siroj05/portfolio/config"
)

type contextKey string

const claimsKey contextKey = "claims"

func JWTauth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var jwtSecret = []byte(config.JWTSecret)
		cookie, err := r.Cookie("session")

		if len(jwtSecret) == 0 {
			log.Println("Jwt secret not set")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		TokenString := cookie.Value

		token, err := jwt.Parse(TokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), claimsKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetClaims(r *http.Request) (jwt.MapClaims, bool) {
	claims, ok := r.Context().Value(claimsKey).(jwt.MapClaims)
	return claims, ok
}
