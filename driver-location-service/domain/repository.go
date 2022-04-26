package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DriverLocationRepository interface {
	CreateDriver(location *DriverLocation) (int64, error)
	UpdateDriver(location *DriverLocation) (int64, error)
	DeleteDriverById(id primitive.ObjectID) (int64, error)
        DriverById(id primitive.ObjectID) (*DriverLocation, error)
        DriverByLocation(location *Location, r float64) (*DriverLocation, error)
	Drivers() (DriverLocations, error)
}
