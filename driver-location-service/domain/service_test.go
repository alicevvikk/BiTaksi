package domain

import (
	"testing"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
)
/*
type MockRepository interface {
        CreateDriver(location *DriverLocation) (int64, error)
        UpdateDriver(location *DriverLocation) (int64, error)
        DeleteDriverById(id primitive.ObjectID) (int64, error)
        DriverById(id primitive.ObjectID) (*DriverLocation, error)
        DriverByLocation(location *Location, r float64) (*DriverLocation, error)
        Drivers() (DriverLocations, error)
}
*/
type mockRepository struct {}

func (mr mockRepository) CreateDriver(location *DriverLocation) (int64, error) {
	return 1, nil
}
func (mr mockRepository) UpdateDriver(location *DriverLocation) (int64, error) {
	return 1, nil

}
func (mr mockRepository) DeleteDriverById(id primitive.ObjectID) (int64, error) {
	return 1, nil
}

func (mr mockRepository) DriverById(id primitive.ObjectID) (*DriverLocation, error) {
	return nil, nil
 }

func (mr mockRepository) DriverByLocation(location *Location, r float64) (*DriverLocation, error) {
	return &driverLocation{
		Id	
	}
}
func (mr mockRepository) Drivers() (DriverLocations, error){
	return nil, nil
}

var service DriverLocationService // This is the service being tested.
var validId primitive.ObjectID // This is a valid ObjID.

func TestMain(m *testing.M) {
        setUp()
        code := m.Run()
        os.Exit(code)
}

func setUp() {
	mock := new(mockRepository)
	service = NewDriverLocationService(mock)
	validId, _ = primitive.ObjectIDFromHex("123456789123456789123456")
//	validAndExistId, _ = primitive.ObjectIDFromHex("62645f75cd9c930bae0d1c62")

}

// Inserted === 2 Updated == 0
func TestCreateDriver(t *testing.T) {
	expectedInsert := int64(2)
	expectedUpdate := int64(0)
        locations := DriverLocations {
                {
                        Location: Location {
                                Type:           "Point",
                                Coordinates:    []float64{12.345678, 25.345678},
                },
                },
                {
                        Location: Location {
                                Type:           "Point",
                                Coordinates:    []float64{12.34567, 25.345678},
                },
                },
}



	ins, upd := service.CreateDriver(locations)
	if ins != expectedInsert || upd != expectedUpdate {
		t.Errorf("EXPECTED -->(%v, %v) GOT --> (%v, %v)", expectedInsert, expectedUpdate,
								   ins, upd)
	}
}


// Inserted === 2 Updated == 1
func TestCreateDriver_2(t *testing.T) {
	expectedInsert := int64(2)
	expectedUpdate := int64(1)
        locations := DriverLocations {
                {
                        Location: Location {
                                Type:           "Point",
                                Coordinates:    []float64{12.345678, 25.345678},
                },
                },
                {
                        Location: Location {
                                Type:           "Point",
                                Coordinates:    []float64{12.34567, 25.345678},
                },
                },
		{
			Location: Location {
				Type:		"Point",
				Coordinates:	[]float64{12.34567, 25.345678},
			},
			Id:	validId,
		},
}



	ins, upd := service.CreateDriver(locations)
	if ins != expectedInsert || upd != expectedUpdate {
		t.Errorf("EXPECTED -->(%v, %v) GOT --> (%v, %v)", expectedInsert, expectedUpdate,
								   ins, upd)
	}
}

// 
// Inserted === 2 Updated == 1
func TestCreateDriver_3(t *testing.T) {
	expectedInsert := int64(2)
	expectedUpdate := int64(1)
        locations := DriverLocations {
                {
                        Location: Location {
                                Type:           "Point",
                                Coordinates:    []float64{12.345678, 25.345678},
                },
                },
                {
                        Location: Location {
                                Type:           "Point",
                                Coordinates:    []float64{12.34567, 25.345678},
                },
                },
		{
			Location: Location {
				Type:		"Point",
				Coordinates:	[]float64{12.34567, 25.345678},
			},
			Id:	validId,
		},
}



	ins, upd := service.CreateDriver(locations)
	if ins != expectedInsert || upd != expectedUpdate {
		t.Errorf("EXPECTED -->(%v, %v) GOT --> (%v, %v)", expectedInsert, expectedUpdate,
								   ins, upd)
	}
}

