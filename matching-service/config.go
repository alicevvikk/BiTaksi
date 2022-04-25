package main

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	Address         string
	ReadTimeout     int64
	WriteTimeout    int64
	ShutdownTimeout int64
	Static          string
}

var (
	config Configuration
)

/*
func initializeConfig() {
	loadConfig()
	file, err := os.OpenFile("matching-service.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}

	logger = log.New(file, "Matching-API", log.LstdFlags)

}
*/

func loadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)

	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}
