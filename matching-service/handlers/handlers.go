package handlers

import (
	"net/http"
	"errors"
	"bytes"

	"github.com/dgrijalva/jwt-go"

	"github.com/alicevvikk/bitaksi/matching-service/data"
	"github.com/alicevvikk/bitaksi/matching-service/logger"

)

var tokenKey = []byte("my_secret_key")
var tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoZW50aWNhdGVkIjp0cnVlfQ.CK7jhYYJ_ULnaO4s_vjy15_6pfFzwI5ns-s4XPvGYyo"

func MatchingHandler(w http.ResponseWriter, r *http.Request) {
	//matchingRequest := &data.MatchingRequest{}
	matchingRequest := new(data.MatchingRequest)
	err := data.FromJSON(r.Body, matchingRequest)

	logger.Log.Println(matchingRequest)
	if err != nil {
		logger.Log.Println("Bad request. from: handlers/handler.MatchingHandler.2")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(matchingRequest.Coordinates) != 2 {
		logger.Log.Println("Bad request. from: handlers/handerl.MatchingHandler")
		err := errors.New("Bad Request")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	buf := new(bytes.Buffer)
	data.ToJSON(buf, matchingRequest)

	req, err := http.NewRequest("POST", "http://localhost:8081/match", buf)
	if err != nil {
		logger.Log.Println("Can't create a request. from: handers/handler.MatchigHandler")
		http.Error(w, "error creating a request", http.StatusInternalServerError)
		return
	}
        req.Header.Add("Authorization", tokenString)
        req.Header.Add("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Log.Println("Can't make a request. from: handlers/handler.MatchingHandler")
		http.Error(w, "error making a request", http.StatusInternalServerError)
		return
	}
	if resp.StatusCode != http.StatusOK {
		logger.Log.Println("No match found")
		http.Error(w, "no match found", resp.StatusCode)
	}

	driverResponse := new(data.DriverResponse)
	data.FromJSON(resp.Body, driverResponse)
	data.ToJSON(w, driverResponse)

}


func MustAuth(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
        return func(w http.ResponseWriter, r*http.Request) {
                token := r.Header.Get("Authorization")
                if token == "" {
			logger.Log.Println("Unauthorized access")
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
			logger.Log.Println("Invalid token!")
                        http.Error(w, err.Error(), http.StatusUnauthorized)
                        return 
                }

                if isAuthenticated := claims["authenticated"]; isAuthenticated != true {
                        logger.Log.Println("Unauthorized access")
			http.Error(w, err.Error(), http.StatusUnauthorized)
                        return
                }

                fn(w, r)
        }

}

