package domain

type DriverLocationService interface {

	CreateDriver(locations Locations) int
	DeleteDriverById(id string) (int64, error)
	DriverById(id string) (DriverLocation, error)
	DriverByLocation(location *Location ) (*ResponseLocation, error)
	Drivers() ([]DriverLocation, error)
}

type driverLocationService struct {
	repo DriverLocationRepository
}

func NewDriverLocationService(r DriverLocationRepository) DriverLocationService {
	return &driverLocationService{
		repo:	r,
	}
}

func (dls *driverLocationService) DeleteDriverById(id string) (int64, error) {
	return dls.repo.DeleteDriverById(id)
}

func (dls *driverLocationService) CreateDriver(locations Locations) int {
	return dls.repo.CreateDriver(locations)
}

func (dls *driverLocationService) DriverById(id string) (DriverLocation, error) {
	return dls.repo.DriverById(id)
}

func (dls *driverLocationService) DriverByLocation(location *Location) (*ResponseLocation, error) {
	return dls.repo.DriverByLocation(location)
}

func (dls *driverLocationService) Drivers() ([]DriverLocation, error) {
	return dls.repo.Drivers()
}
