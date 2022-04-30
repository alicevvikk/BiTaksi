package repository 

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/alicevvikk/bitaksi/driver-location-service/logger"

	"go.mongodb.org/mongo-driver/bson"
)

var locations []interface{}

func parseFloat(val string) (float64, error){
	valFloat, err := strconv.ParseFloat(val, 64)
	if err!= nil {
		return 0, err
	}
	return valFloat, nil
}

func loadCoordinates() {
	count := 0
	f, err := os.Open("repository/coordinates.csv")
	if err != nil {
		logger.Fatal(err)
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		logger.Fatal("can't load coordinates into memory. from: repo.mongo.loadCoordinates")
	}

	lines = lines[1:]
	for _, line := range lines {

		lat, err1 := parseFloat(line[0])
		lon, err2 := parseFloat(line[1])

		if err1 != nil || err2 != nil {
			continue
		}

		location := bson.D{{"location", bson.D{
			{"type", "Point"},
			{"coordinates", []float64{lon, lat}},
		}}}

		locations = append(locations, location)
		count ++
	}
	logger.Infof("%d coordinates has been loaded into memory", count)
}






