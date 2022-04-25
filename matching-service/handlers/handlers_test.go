package handlers

import (
	"net/http"
	"testing"
	"github.com/alicevvikk/bitaksi/matching-service/data"
	"encoding/json"
	"bytes"
)

var tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoZW50aWNhdGVkIjp0cnVlfQ.CK7jhYYJ_ULnaO4s_vjy15_6pfFzwI5ns-s4XPvGYyo"

func TestMatchHandlerAuthentication200(t *testing.T) {
	expectedStatusCode := http.StatusOK // 200

	model := data.MatchingRequest{
		Type:		"Point",
		Coordinates:	[]float64{21.12345, 22.12345},
	}
	var newBuffer = &bytes.Buffer{}
	err := model.ToJSON(newBuffer)

	if err != nil {
		t.Fatal(err)
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8080/match/",
				   newBuffer)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", tokenString)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("Cant make a request", err)
	}

	if resp.StatusCode != expectedStatusCode {
		t.Errorf("got %d expected %d", resp.StatusCode, expectedStatusCode)
	}

}

func TestMatchHandlerAuthentication400(t *testing.T) {
	expectedStatusCode := http.StatusBadRequest  // 400

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8080/match/", nil)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", tokenString)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("Cant make a request", err)
	}

	if resp.StatusCode != expectedStatusCode {
		t.Errorf("got %d expected %d", resp.StatusCode, expectedStatusCode)
	}

}

func TestMatchHandlerAuthentication400_2(t *testing.T) {
	expectedStatusCode := http.StatusBadRequest  // 400

	model := data.MatchingRequest{
		Type:		"Point",
	}
	modelBuffer := &bytes.Buffer{}
	err := model.ToJSON(modelBuffer)

	if err != nil {
		t.Fatal(err)
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8080/match/",
				    modelBuffer)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", tokenString)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("Cant make a request", err)
	}

	if resp.StatusCode != expectedStatusCode {
		t.Errorf("got %d expected %d", resp.StatusCode, expectedStatusCode)
	}

}

func TestMatchHandlerAuthentication401(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized  // 401

	model := data.MatchingRequest{
		Type:		"Point",
		Coordinates:	[]float64{21.123456, 22.123456},
	}
	modelBuffer := &bytes.Buffer{}
	err := model.ToJSON(modelBuffer)
	if err != nil {
		t.Fatal(err)
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8080/match/",
				    modelBuffer)

	if err != nil {
		t.Fatal(err)
	}

	//req.Header.Add("Authorization", tokenString)
	//req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("Cant make a request", err)
	}

	if resp.StatusCode != expectedStatusCode {
		t.Errorf("got %d expected %d", resp.StatusCode, expectedStatusCode)
	}

}

func TestMatchHandlerAuthentication401_2(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized  // 401

	model := data.MatchingRequest{
		Type:		"Point",
		Coordinates:	[]float64{24.123456, 222.123456},
	}
	modelBuffer := &bytes.Buffer{}
	err := model.ToJSON(modelBuffer)
	if err != nil {
		t.Fatal(err)
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8080/match/",
				    modelBuffer)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("lalalaland", tokenString)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("Cant make a request", err)
	}

	if resp.StatusCode != expectedStatusCode {
		t.Errorf("got %d expected %d", resp.StatusCode, expectedStatusCode)
	}

}

func TestMatchHandlerAuthentication401_3(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized  // 401

	model := data.MatchingRequest{
		Type:		"Point",
		Coordinates:	[]float64{22.179384, 29.222123},
	}
	modelBuffer := &bytes.Buffer{}
	err := model.ToJSON(modelBuffer)

	if err != nil {
		t.Fatal(err)
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8080/match/",
				    modelBuffer)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "wrong token")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("Cant make a request", err)
	}

	if resp.StatusCode != expectedStatusCode {
		t.Errorf("got %d expected %d", resp.StatusCode, expectedStatusCode)
	}

}
