package domain

import (
	"github.com/alicevvikk/bitaksi/driver-location-service/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Radius value for specifying the maximum distance
// that we can match a driver. IN METERS. 
const (
	radius = 3000.0
)

// Group of methods to serve the business logic.
type DriverLocationService interface {

	CreateDriver(locations DriverLocations) (int64, int64)
	DeleteDriverById(id string) (int64, error)
	DriverById(id string) (*DriverLocation, error)
	DriverByLocation(location *Location ) (*ResponseLocation, error)
	Drivers() (DriverLocations, error)
}

//Service that implements the DriverLocationService interface.
type driverLocationService struct {
	repo DriverLocationRepository
}

// Constructor for DriverLocationService
func NewDriverLocationService(r DriverLocationRepository) DriverLocationService {
	return &driverLocationService{
		repo:	r,
	}
}
// Takes an id and converts it to 'ObjectID' type.
// If given id is not convertable to ObjectID then returns
// 0 as delete count. 
func (dls *driverLocationService) DeleteDriverById(id string) (int64, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}

	return dls.repo.DeleteDriverById(objId)
}


// Takes 'domain/model.DriverLocations' as paramter and decides 
// whether it is an update or create opearation for each driver location.
// Then calls the proper operation on repo(CREATE or UPDATE).
// returns INSERT and UPDATE counts.
func (dls *driverLocationService) CreateDriver(locations DriverLocations) (int64, int64) {
	inserted := int64(0)
	updated := int64(0)

	for _, location := range locations {
		if location.Id.IsZero() {
			res, err := dls.repo.CreateDriver(location)
			if err == nil && res != 0{
				inserted++
			}
			continue
		}
		res, err := dls.repo.UpdateDriver(location)
		if err == nil && res != 0{
			updated++
		}

	}
	return inserted, updated
}

// Takes an id and converts it to 'ObjectID' type.
// If given id is not convertable to ObjectID, then
// returns nil and error.
func (dls *driverLocationService) DriverById(id string) (*DriverLocation, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return dls.repo.DriverById(objId)
}


// Takes a location and calls repository adapter with predetermined
// business radius.
func (dls *driverLocationService) DriverByLocation(location *Location) (*ResponseLocation, error) {
	driverLocation, err := dls.repo.DriverByLocation(location, radius)
	if err != nil {
		return nil, err
	}
	distance := utils.CalculateDistance(
		location.Coordinates,
		driverLocation.Location.Coordinates)

	response := &ResponseLocation{
		DriverLocation:	*driverLocation,
		Distance:	distance,
	}
	return response, nil
}
// Gets all drivers from repository.
func (dls *driverLocationService) Drivers() (DriverLocations, error) {
	return dls.repo.Drivers()
}
