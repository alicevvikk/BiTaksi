package api


import (
	"net/http"
	"testing"
	"os"
	"bytes"

	"github.com/alicevvikk/bitaksi/driver-location-service/domain"
	"github.com/alicevvikk/bitaksi/driver-location-service/utils"

	prim"go.mongodb.org/mongo-driver/bson/primitive"
)

var client *http.Client

var tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoZW50aWNhdGVkIjp0cnVlfQ.CK7jhYYJ_ULnaO4s_vjy15_6pfFzwI5ns-s4XPvGYyo"


type CreateResponse struct {
	TotalReceived	int
	Inserted	int
	Updated		int

}


func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	client = &http.Client{}
	//TODO 
}

//Creates one
func TestCreateDriver_201(t * testing.T) {
	expectedStatusCode := http.StatusCreated
	method := "POST"
	URL := "http://localhost:8081/driver"
	expectedInsert := 1 


	models := domain.DriverLocations {
		{
			Location: domain.Location {
				Type:		"Point",
				Coordinates:	[]float64{12.345678, 25.345678},
		},
		},
	}

	buf := new(bytes.Buffer)
	utils.ToJSON(buf, models)

	req, err := http.NewRequest(method, URL, buf)
	if err != nil {
		t.Error(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != expectedStatusCode {
		t.Errorf("GOT --> %v EXPECTED --> %v", resp.StatusCode, expectedStatusCode)
	}
	response := new(CreateResponse) 
	err = utils.FromJSON(resp.Body, response)

	if err != nil {
		t.Error(err)
	}

	if response.Inserted != expectedInsert {
		t.Errorf("EXPECTED INSERT --> %v GOT --> %v", expectedInsert, response.Inserted)
	}

}


//Creates two 
func TestCreateDriver_201_2(t * testing.T) {
	expectedStatusCode := http.StatusCreated
	method := "POST"
	URL := "http://localhost:8081/driver"
	expectedInsert := 2


	models := domain.DriverLocations {
		{
			Location: domain.Location {
				Type:		"Point",
				Coordinates:	[]float64{12.345678, 25.345678},
		},
		},
		{
			Location: domain.Location {
				Type:		"Point",
				Coordinates:	[]float64{12.34567, 25.345678},
		},
		},
}

	buf := new(bytes.Buffer)
	utils.ToJSON(buf, models)

	req, err := http.NewRequest(method, URL, buf)
	if err != nil {
		t.Error(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != expectedStatusCode {
		t.Errorf("GOT --> %v EXPECTED --> %v", resp.StatusCode, expectedStatusCode)
	}
	response := new(CreateResponse) 
	err = utils.FromJSON(resp.Body, response)
	
	if err != nil {
		t.Error(err)
	}

	if response.Inserted != expectedInsert {
		t.Errorf("EXPECTED INSERT --> %v GOT --> %v", expectedInsert, response.Inserted)
	}

}


//Creates two, updates one 
func TestCreateDriver_201_3(t * testing.T) {
	expectedStatusCode := http.StatusCreated
	method := "POST"
	URL := "http://localhost:8081/driver"
	expectedInsert := 2
	expectedUpdate := 1

	id, _:= prim.ObjectIDFromHex("62645f75cd9c930bae0d1c5c")
	models := domain.DriverLocations {
		{
			Location: domain.Location {
				Type:		"Point",
				Coordinates:	[]float64{12.345678, 25.345678},
			},
		},
		{
			Location: domain.Location {
				Type:		"Point",
				Coordinates:	[]float64{12.34567, 25.345678},
			},
		},
		{
			Location: domain.Location {
				Type:		"Point",
				Coordinates:	[]float64{12.34567, 22.345678},
			},
			Id:	id,
		},
}

	buf := new(bytes.Buffer)
	utils.ToJSON(buf, models)

	req, err := http.NewRequest(method, URL, buf)
	if err != nil {
		t.Error(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != expectedStatusCode {
		t.Errorf("GOT --> %v EXPECTED --> %v", resp.StatusCode, expectedStatusCode)
	}

	response := new(CreateResponse)
	err = utils.FromJSON(resp.Body, response)

	if err != nil {
		t.Error(err)
	}

	if response.Inserted != expectedInsert || response.Updated != expectedUpdate {
		t.Errorf("GOT --> (%v, %v) EXPECTED --> (%v, %v)",
			response.Inserted, response.Updated, expectedInsert, expectedUpdate)
	}
}

// If invalid ID, create new
// 
func TestCreateDriver_201_4(t * testing.T) {
	expectedStatusCode := http.StatusCreated
	method := "POST"
	URL := "http://localhost:8081/driver"
	expectedInsert := 3
	expectedUpdate := 0

	// invalid id, so id gets zero value
	id, _:= prim.ObjectIDFromHex("123456")
	models := domain.DriverLocations {
		{
			Location: domain.Location {
				Type:		"Point",
				Coordinates:	[]float64{12.345678, 25.345678},
			},
		},
		{
			Location: domain.Location {
				Type:		"Point",
				Coordinates:	[]float64{12.34567, 25.345678},
			},
		},
		{
			Location: domain.Location {
				Type:		"Point",
				Coordinates:	[]float64{12.34567, 25.345678},
			},
			Id:	id,
		},
}

	buf := new(bytes.Buffer)
	utils.ToJSON(buf, models)

	req, err := http.NewRequest(method, URL, buf)
	if err != nil {
		t.Error(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != expectedStatusCode {
		t.Errorf("GOT --> %v EXPECTED --> %v", resp.StatusCode, expectedStatusCode)
	}

	response := new(CreateResponse)
	err = utils.FromJSON(resp.Body, response)

	if err != nil {
		t.Error(err)
	}

	if response.Inserted != expectedInsert || response.Updated != expectedUpdate {
		t.Errorf("GOT --> (%v, %v) EXPECTED --> (%v, %v)",
			response.Inserted, response.Updated, expectedInsert, expectedUpdate)
	}
}

func TestDriverByLocation_200(t *testing.T) {
	expectedStatusCode := http.StatusOK
	method := "POST"
	url := "http://localhost:8081/match"

	model := domain.Location{
		Type:		"Point",
		Coordinates:	[]float64{29.0390200, 42.0000001},
	}
	buf := &bytes.Buffer{}
	utils.ToJSON(buf, model)

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		t.Error(err)
	}
	req.Header.Add("Authorization", tokenString)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != expectedStatusCode {
		t.Errorf("GOT --> %v EXPECTED --> %v", resp.StatusCode, expectedStatusCode)
	}
}
//Missing coordinate.
func TestDriverByLocation_404(t *testing.T) {
	expectedStatusCode := http.StatusNotFound
	method := "POST"
	URL := "http://localhost:8081/match"

	model := domain.Location{
		Type:		"Point",
		Coordinates:	[]float64{42.0000001},
	}
	buf := &bytes.Buffer{}
	utils.ToJSON(buf, model)

	req, err := http.NewRequest(method, URL, buf)
	if err != nil {
		t.Error(err)
	}
	req.Header.Add("Authorization", tokenString)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != expectedStatusCode {
		t.Errorf("GOT --> %v EXPECTED --> %v", resp.StatusCode, expectedStatusCode)
	}

}

// send nil coordinates
func TestDriverByLocation_404_2(t *testing.T) {
	expectedStatusCode := http.StatusNotFound
	method := "POST"
	URL := "http://localhost:8081/match"

	model := domain.Location{
		Type:	"Point",
	}

	buf := &bytes.Buffer{}
	utils.ToJSON(buf, model)

	req, err := http.NewRequest(method, URL, buf)
	if err != nil {
		t.Error(err)
	}

	req.Header.Add("Authorization", tokenString)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != expectedStatusCode {
		t.Errorf("GOT --> %v EXPECTED --> %v", resp.StatusCode, expectedStatusCode)
	}

}

//valid coordinates, but not matched with a driver.
func TestDriverByLocation_404_3(t *testing.T) {
	expectedStatusCode := http.StatusNotFound
	method := "POST"
	URL := "http://localhost:8081/match"

	model := domain.Location{
		Type:		"Point",
		Coordinates:	[]float64{129.0390200, 42.1000001},
	}

	buf := &bytes.Buffer{}
	utils.ToJSON(buf, model)

	req, err := http.NewRequest(method, URL, buf)
	if err != nil {
		t.Error(err)
	}

	req.Header.Add("Authorization", tokenString)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != expectedStatusCode {
		t.Errorf("GOT --> %v EXPECTED --> %v", resp.StatusCode, expectedStatusCode)
	}

}


func TestDriverByLocation_401(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized

	model := domain.Location{
		Type:		"Point",
		Coordinates:	[]float64{29.0390200, 42.0000001},
	}
	buf := &bytes.Buffer{}
	utils.ToJSON(buf, model)

	req, err := http.NewRequest("POST", "http://localhost:8081/match", buf)
	if err != nil {
		t.Error(err)
	}
//	req.Header.Add("Authorization", tokenString)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != expectedStatusCode {
		t.Errorf("GOT --> %v EXPECTED --> %v", resp.StatusCode, expectedStatusCode)
	}
}

func TestDriverByLocation_401_2(t *testing.T) {
        expectedStatusCode := http.StatusUnauthorized
        method := "POST"
        URL := "http://localhost:8081/match"

        model := domain.Location{
                Type:           "Point",
                Coordinates:    []float64{29.0390200, 42.0000001},
        }

        buf := &bytes.Buffer{}
        utils.ToJSON(buf, model)
        req, err := http.NewRequest(method, URL, buf)
        if err != nil {
                t.Error(err)
        }

        req.Header.Add("Authorization", "WRONG_KEY")
        req.Header.Add("Content-Type", "application/json")

        resp, err := client.Do(req)
        if err != nil {
                t.Error(err)
        }

        if resp.StatusCode != expectedStatusCode {
                t.Errorf("GOT --> %v EXPECTED --> %v", resp.StatusCode, expectedStatusCode)
        }

}

//fasfasfasfasfas
func TestDriverByLocation_405(t *testing.T) {
	expectedStatusCode := http.StatusMethodNotAllowed
	method := "GET"
	URL := "http://localhost:8081/match"

	req, err := http.NewRequest(method, URL, nil)
	if err != nil {
		t.Error(err)
	}

	resp, err := client.Do(req)
	if  err != nil {
		t.Error(err)
	}

	if resp.StatusCode != expectedStatusCode {
		t.Errorf("GOT --> %v EXPECTED --> %v", resp.StatusCode, expectedStatusCode)
	}

}

//correct id
func TestDriverById_200(t *testing.T) {
	expectedStatusCode := http.StatusOK
	method := "GET"
	id := "62645f75cd9c930bae0d1c60"
	URL := "http://localhost:8081/driver/"

	req, err := http.NewRequest(method, URL + id, nil)
	if err != nil {
		t.Error(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != expectedStatusCode {
		t.Errorf("GOT --> %v EXPECTED --> %v", resp.StatusCode, expectedStatusCode)
	}


}

//wrong id
func TestDriverById_404(t *testing.T) {
	expectedStatusCode := http.StatusNotFound
	method := "GET"
	id := "WRONG_ID"
	URL := "http://localhost:8081/driver/"

	req, err := http.NewRequest(method, URL + id, nil)
	if err != nil {
		t.Error(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != expectedStatusCode {
		t.Errorf("GOT --> %v EXPECTED --> %v", resp.StatusCode, expectedStatusCode)
	}


}

//delete driver with correct id
func TestDeleteDriverById_200(t *testing.T) {
	expectedStatusCode := http.StatusOK
	method := "DELETE"
	id := "62645f75cd9c930bae0d1c63"
	URL := "http://localhost:8081/driver/"

	req, err := http.NewRequest(method, URL + id, nil)
	if err != nil {
		t.Error(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != expectedStatusCode {
		t.Errorf("GOT --> %v EXPECTED --> %v", resp.StatusCode, expectedStatusCode)
	}
}

//wrong id
func TestDeleteDriverById_404(t *testing.T) {
	expectedStatusCode := http.StatusNotFound
	method := "DELETE"
	id := "WRONG_ID"
	URL := "http://localhost:8081/driver/"

	req, err := http.NewRequest(method, URL + id, nil)
	if err != nil {
		t.Error(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != expectedStatusCode {
		t.Errorf("GOT --> %v EXPECTED --> %v", resp.StatusCode, expectedStatusCode)
	}
}
