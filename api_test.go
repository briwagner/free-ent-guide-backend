package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type requestTest struct {
	url      string
	expected string
	code     int
}

func TestGetMovies(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/movies", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetMovies)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Hander returned wrong status. Got %v, Want %v", status, http.StatusOK)
	}
}

func TestDiscoverMovies(t *testing.T) {
	requests := []requestTest{
		{"/v1/discover", "Must pass a date", 406},
		{"/v1/discover?date=2019-05-05", `{"status_code":7,"status_message":"Invalid API key: You must be granted a valid key.","success":false}`, 200},
	}

	for _, tt := range requests {
		req, err := http.NewRequest("GET", tt.url, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(DiscoverMovies)

		handler.ServeHTTP(rr, req)

		expected := tt.code
		if status := rr.Code; status != expected {
			t.Errorf("Hander returned wrong status. Got %v, Want %v", status, expected)
		}

		expectedBody := tt.expected
		if rr.Body.String() != expectedBody {
			t.Errorf("Handler returned wrong body. Got %v, Want %v", rr.Body.String(), expectedBody)
		}
	}
}

func TestGetTvMovies(t *testing.T) {
	requests := []requestTest{
		{"/v1/tv-movies", "Must pass a date", 406},
		{"/v1/tv-movies?date=2019-05-28", `{"status_code":7,"status_message":"Invalid API key: You must be granted a valid key.","success":false}`, 200},
	}

	for _, tt := range requests {
		req, err := http.NewRequest("GET", tt.url, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(GetTvMovies)

		handler.ServeHTTP(rr, req)

		expected := tt.code
		if status := rr.Code; status != expected {
			t.Errorf("Hander returned wrong status. Got %v, Want %v", status, expected)
		}

		expectedBody := tt.expected
		if rr.Body.String() != expectedBody {
			t.Errorf("Handler returned wrong body. Got %v, Want %v", rr.Body.String(), expectedBody)
		}
	}
}

func TestGetTMSReq(t *testing.T) {
	// get creds
	// mock loc
	// make REAL request
	// test response
}
