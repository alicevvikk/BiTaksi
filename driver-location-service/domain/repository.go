package domain


type DriverLocationRepository interface {
	CreateDriver(locations Locations) int
	DeleteDriverById(id string) (int64, error)
        DriverById(id string) (DriverLocation, error)
        DriverByLocation(location *Location) (*ResponseLocation, error)
	Drivers() ([]DriverLocation, error)
}
