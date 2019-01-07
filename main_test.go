package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var api API

func TestMain(m *testing.M) {
	api = API{}
	api.Init()

	code := m.Run()

	os.Exit(code)
}

func TestGetDeviceStatsNonExistent(t *testing.T) {
	req, _ := http.NewRequest("GET", "/devices/1c/stats", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)
}

func TestGetDeviceStats(t *testing.T) {
	req, _ := http.NewRequest("GET", "/devices/1/stats", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	api.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
