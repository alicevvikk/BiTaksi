package data

import (
	"encoding/json"
	"io"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


// This type is response to a end-user.
type DriverResponse struct {
	DriverLocation	DriverLocation	`json:"driverLocation"`
	Distance	float64		`json:"distance"`
}

// This type represents a driver.
type DriverLocation struct {
	Id		primitive.ObjectID `json:"id" bson:"_id"`
	Type		string	           `json:"type" bson:"type"`
	Location	MatchingRequest	   `json: location bson:"location"`
}


// This is the request initiated by the end user
// for matching with a driver.
type MatchingRequest struct{
//	Type		string	   `json:"type" bson:"type"`
	Coordinates	[]float64  `json:"coordinates" bson"coordinates"`
}

// Takes any data type, encodes data to JSON,
// then writes into 'w'.
func ToJSON(w io.Writer, data interface{}) error {
	newEncoder := json.NewEncoder(w)
	return newEncoder.Encode(data)
}

// Reads from 'io.Reader', converts the JSON
// data and writes into 'data'
func FromJSON(r io.Reader, data interface{}) error {
	newDecoder := json.NewDecoder(r)
	return newDecoder.Decode(data)
}

