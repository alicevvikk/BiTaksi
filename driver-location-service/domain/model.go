package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"encoding/json"
	"io"
)

type LocationType interface {

}

type ResponseLocation struct {
	DriverLocation	DriverLocation	`json:"driverLocation"`
	Distance	float64		`json:"distance"`
}

type DriverLocation struct {
	Id		primitive.ObjectID `json:"id" bson:"_id"`
	Location	Location	   `json: location bson:"location"`
}

type Location struct {
	Type		string	  `json:"type" bson:"type"`
	Coordinates	[]float64 `json:"coordinates" bson:"coordinates"`
}

type Locations []*DriverLocation

/*
func (l *Locations)FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)


func (l* Location) ToJSON(w io.Writer) error {
	newEncoder := json.NewEncoder(w)
	return newEncoder.Encode(l)
}

func (l* Location)FromJSON(r io.Reader) error {
	newDecoder := json.NewDecoder(r)
	return newDecoder.Decode(l)
}

func (l* DriverLocation) ToJSON(w io.Writer) error {
	newEncoder := json.NewEncoder(w)
	return newEncoder.Encode(l)
}

func (l* DriverLocation)FromJSON(r io.Reader) error {
	newDecoder := json.NewDecoder(r)
	return newDecoder.Decode(l)
}
*/

func ToJSON (w io.Writer, l LocationType) error {
	newEncoder := json.NewEncoder(w)
        return newEncoder.Encode(l)
}

func FromJSON(r io.Reader, l LocationType) error {
        newDecoder := json.NewDecoder(r)
        return newDecoder.Decode(l)
}
