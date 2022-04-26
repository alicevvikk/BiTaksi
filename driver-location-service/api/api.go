package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/alicevvikk/bitaksi/driver-location-service/logger"
	"github.com/alicevvikk/bitaksi/driver-location-service/domain"
	"github.com/alicevvikk/bitaksi/driver-location-service/utils"
)



// Group of methods to serve HTTP requests.
type Handler interface {
	Get(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	Match(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
}

type handler struct {
	service domain.DriverLocationService
}

func NewHandler(driverLocationService domain.DriverLocationService) Handler {
	return &handler{service:driverLocationService}
}

//#Description: Get handles GET requests and returns all drivers or a
//specific driver if ID is given in the url
//#route: /driver OR /driver/{id}
//#method: GET
//#if SUCCESS --> returns |202, domain/model.driverLocation
//#if FAILURE --> Returns  404 
func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	logger.Info("GetDriver(s).")
	w.Header().Set("Content-type", "application/json")
	id := chi.URLParam(r, "id")
	if id == "" {
		driverLocations, err:= h.service.Drivers()
		if err != nil {
			logger.Error("Can't get drivers")
			http.Error(w, "Can't get drivers", http.StatusNotFound)
		}
		err = utils.ToJSON(w, driverLocations)
		return
	}

	driverLocation, err := h.service.DriverById(id)
	if err != nil {
		logger.Error("Driver not found.")
		http.Error(w, "NotFound", http.StatusNotFound)
	}

	err = utils.ToJSON(w, driverLocation)
	return
}

//#Description: Delete handler calls the service that handles DELETE requests and deletes one user
//with the given ID specified  in the url.
//#route /driver/{id} 
//#method DELETE
//# if SUCCESS --> returns 202
//# if FAILURE --> returns 404
func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	logger.Info("Delete driver.")
	id := chi.URLParam(r, "id")
	if id == "" {
		logger.Error("Empty ID. FROM: handler.Delete")
		http.Error(w, "", http.StatusNoContent)
		return
	}

	_, err := h.service.DeleteDriverById(id)
	if err != nil {
		logger.Error("Can't delete. FROM: handler.Delete")
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	logger.Infof("User with the id: %d deleted", id)
}

//#Description: Calls the service that creates or updates one or many users.
//Request body sent to this method must be given in 'domain/model.Locations' type.
//#route: /driver
//#method: POST
//#if SUCCESS --> returns 201, domain/model.createResponse 
//#if FAILURE --> returns 400, 500
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	logger.Info("Create or update driver(s).")
	locations := domain.DriverLocations{}
	err := utils.FromJSON(r.Body, &locations)
	log.Println(*locations[0])
	if err != nil {
		logger.Error("server error. FROM: handler.Create")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	inserted, updated := h.service.CreateDriver(locations)

	if inserted + updated == 0  && len(locations) > 0{
		logger.Error("Bad request. FROM: handler.Create")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := domain.CreateResponse {
		TotalReceived:	int64(len(locations)),
		Inserted:	inserted,
		Updated:	updated,
	  }

	err = utils.ToJSON(w, response)
	if err != nil {
		logger.Error("Server Error. FROM: handler.Create")
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	logger.Infof("%d documents created %d documents updated", inserted, updated)
	w.WriteHeader(http.StatusCreated)

}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	return
}
//#Description: Calls the service method for finding the
//nearest driver with the 'domain/model.userLocation' type as parameter.
//This endpoint is protected by 'MustAuth' function below.
//#route: /match
//#method: POST 
//#if SUCCESS --> returns 200, 'domain/model.ResponseLocation'
//#if FAILURE --> returns 400, 404, 401(protected by 'middleware/auth.MustAuth')
func (h *handler) Match(w http.ResponseWriter, r *http.Request) {
	logger.Info("Match request")
	userLocation := new(domain.Location)
	err := utils.FromJSON(r.Body, userLocation)
	if err != nil {
		logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("LOCATON_RECEIVED: ", userLocation)
	result, err := h.service.DriverByLocation(userLocation)
	log.Println("RES: ", result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	err = utils.ToJSON(w, result)
	if err != nil {
		log.Println("here")
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

