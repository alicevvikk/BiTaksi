package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LocationType interface {}

// This type is used as a repsonse for the 
// requests to 'api/api.Create' endpoint.
type CreateResponse struct {
        TotalReceived int64 `json:"totelReceived"`
        Inserted      int64 `json:"inserted"`
        Updated       int64 `json:"updated"`
}
// This type is used as a response for the
// requests to 'api/api.Match' endpoint.
type ResponseLocation struct {
	DriverLocation	DriverLocation	`json:"driverLocation"`
	Distance	float64		`json:"distance"`
}
// This type is representation of a document
// in mongodb database.
type DriverLocation struct {
	Id		primitive.ObjectID `json:"id" bson:"_id"`
	Location	Location	   `json: location bson:"location"`
}
// This type is used as a representation
// of a GeoJSON object.
type Location struct {
	Type		string	  `json:"type" bson:"type"`
	Coordinates	[]float64 `json:"coordinates" bson:"coordinates"`
}


// This type holds pointers to the DriverLocation
// Instead of retrieving
type DriverLocations []*DriverLocation



