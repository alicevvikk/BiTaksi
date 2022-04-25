package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	mr"github.com/alicevvikk/bitaksi/driver-location-service/repository/mongo"
	"github.com/alicevvikk/bitaksi/driver-location-service/domain"
	"github.com/alicevvikk/bitaksi/driver-location-service/api"
)



func main() {

	mongoRepository, err := mr.NewMongoRepository("bitaksi-db", 50)

	if err != nil {
		log.Fatal("Error creating repository")
	}
	//mr.ImportInitialData(mongoRepository)

	service := domain.NewDriverLocationService(mongoRepository)
	handler := api.NewHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/driver", handler.Create)
	r.Get("/driver", handler.Get)
	r.Get("/driver/{id}", handler.Get)
	r.Delete("/driver/{id}", handler.Delete)
	r.Post("/match", api.MustAuth(handler.Match))

	server := http.Server{
		Addr:		":8081",
		Handler:	r,
	}

	log.Println("listening..")
	log.Fatal(server.ListenAndServe())
}


