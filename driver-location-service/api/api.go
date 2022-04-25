package api

import (
	"log"
	"net/http"
	"encoding/json"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"

	"github.com/alicevvikk/bitaksi/driver-location-service/domain"
)

var tokenKey = []byte("my_secret_key")

type Handler interface {
	Get(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	Match(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
}

type handler struct {
	service domain.DriverLocationService
}

func NewHandler(driverLocationService domain.DriverLocationService) Handler {
	return &handler{service:driverLocationService}
}


func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	id := chi.URLParam(r, "id")
	if id == "" {
		driverLocations, err:= h.service.Drivers()
		if err != nil {
			http.Error(w, "Can't get drivers", http.StatusNotFound)
		}
		err = json.NewEncoder(w).Encode(&driverLocations)
		return
	}

	driverLocation, err := h.service.DriverById(id)
	if err != nil {
		http.Error(w, "NotFound", http.StatusNotFound)
	}

	err = json.NewEncoder(w).Encode(&driverLocation)
	return
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	log.Println("TARGET ID ", id)
	if id == "" {
		http.Error(w, "", http.StatusNotFound)
	}

	count, err := h.service.DeleteDriverById(id)
	if err != nil || count == 0{
		http.Error(w, "Not Foun", http.StatusNotFound)
	}

}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	locations := domain.Locations{}
	err := domain.FromJSON(r.Body, &locations)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	h.service.CreateDriver(locations)
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *handler) Match(w http.ResponseWriter, r *http.Request) {
	userLocation := new(domain.Location)
	//userLocation := &domain.Location{}
	err := domain.FromJSON(r.Body, userLocation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("LOCA:	", userLocation)
	result, err := h.service.DriverByLocation(userLocation)
	log.Println("RES: ", result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	err = domain.ToJSON(w, result)
	if err != nil {
		log.Println("here")
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

func MustAuth(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r*http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}

		claims := jwt.MapClaims{}
		tkn, err := jwt.ParseWithClaims(
				token,
				claims,
				func(token *jwt.Token) (interface{}, error) {
			return tokenKey, nil
		})

		if err != nil || !tkn.Valid {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if isAuthenticated := claims["authenticated"]; isAuthenticated != true {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		fn(w, r)
	}

}
