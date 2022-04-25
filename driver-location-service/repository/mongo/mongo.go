package repository 

import (
	"encoding/csv"
	"context"
	"log"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var locations []interface{}

func ImportInitialData(mr *mongoRepository) {
	log.Println("Starting to import initial data.. from: repo.mongo.ImportInitialData")
	loadCoordinates()
	createLocationIndex(mr)

	coll := mr.client.Database(mr.db).Collection("driver-locations")

	ctx, cancel := context.WithTimeout(context.Background(), mr.timeout)
	defer cancel()

	_, err := coll.InsertMany(ctx, locations)
	if err != nil {
		log.Fatal("error. from repo.mongo.ImportInitialData")
	}
	log.Println("Initial data imporeted successfully.. from: repo.mongo.ImportInitialData")
}

func parseFloat(val string) float64{
	valFloat, err := strconv.ParseFloat(val, 64)
	if err!= nil {
		log.Fatal("from: repo.mongo.parseFloat")
	}
	return valFloat
}

func loadCoordinates() {
	f, err := os.Open("repository/coordinates.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatal("can't load coordinates into memory. from: repo.mongo.loadCoordinates")
	}

	lines = lines[1:]
	for _, line := range lines {

		lat := parseFloat(line[0])
		lon := parseFloat(line[1])

		location := bson.D{{"location", bson.D{
			{"type", "Point"},
			{"coordinates", []float64{lon, lat}},
		}}}

		locations = append(locations, location)
	}
}

func createLocationIndex(mr *mongoRepository) {
	log.Println("Starting to create location index. from: repo.mongo.createLocationIndex")
	ctx, cancel := context.WithTimeout(context.Background(), mr.timeout)
	defer cancel()

	model := mongo.IndexModel{
		Keys: bson.D{
			{"location", "2dsphere"},
		},
	}

	coll := mr.client.Database(mr.db).Collection("driver-locations")

	var indexView mongo.IndexView
	indexView = coll.Indexes()
	_, err := indexView.CreateOne(
		ctx,
		model,
		nil,
	)

	if err != nil {
		log.Fatal("Can't create the index. from: repo.mongo.createLocationIndex")
	}
	log.Println("Index created successfully from: repo.mongo.createLocationInex")
}




