package middleware

import (
	"net/http"
	
	"github.com/dgrijalva/jwt-go"
	"github.com/alicevvikk/bitaksi/driver-location-service/logger"
)

var tokenKey = []byte("my_secret_key")

//#Description: Takes a 'http.HandlerFunc' endpoint as parameter and acts as a middleware.
func MustAuth(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r*http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			logger.Error("Unauthorized access.")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		claims := jwt.MapClaims{}
		tkn, err := jwt.ParseWithClaims(
				token,
				claims,
				func(token *jwt.Token) (interface{}, error) {
			return tokenKey, nil
		})

		if err != nil || !tkn.Valid {
			logger.Error("Unauthorized access.")
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if isAuthenticated := claims["authenticated"]; isAuthenticated != true {
			logger.Error("Unauthorized access.")
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		fn(w, r)
	}

}
