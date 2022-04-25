package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/alicevvikk/bitaksi/matching-service/data"
	"github.com/alicevvikk/bitaksi/matching-service/logger"
	"net/http"
	"errors"
)

var tokenKey = []byte("my_secret_key")

func MatchingHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Can't get the jwt key", http.StatusUnauthorized)
		return
	}

	claims := jwt.MapClaims{}
	tkn, err := jwt.ParseWithClaims(token,
					claims,
					func(token *jwt.Token) (interface{}, error) {
		return tokenKey, nil
	})

	if err != nil || !tkn.Valid{
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if isAuthenticated := claims["authenticated"]; isAuthenticated != true {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	matchingRequest := &data.MatchingRequest{}
	err = matchingRequest.FromJSON(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(matchingRequest.Coordinates) != 2 {
		err := errors.New("Bad Request")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

}
