package main

import (
	h"github.com/alicevvikk/bitaksi/matching-service/handlers"
	"github.com/alicevvikk/bitaksi/matching-service/logger"
	"net/http"
	"time"
	"os"
	"os/signal"
	"context"
)

func main() {

	loadConfig()
	
	mux := http.NewServeMux()
	mux.HandleFunc("/match/", h.MustAuth(h.MatchingHandler))

	server := &http.Server{
		Addr:         config.Address,
		ReadTimeout:  time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout: time.Duration(config.WriteTimeout * int64(time.Second)),
		Handler:      mux,
	}

	go func() {
		logger.Log.Println("Starting server on port", config.Address)

		err := server.ListenAndServe()
		if err != nil {
			logger.Log.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	// Block until a signal is received.
	sig := <-signalChannel
	logger.Log.Println("Shutting down the server. Signal : ", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)


}
