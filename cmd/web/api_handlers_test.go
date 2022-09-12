package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type requestTest struct {
	url      string
	expected string
	code     int
}

func TestGetMovies(t *testing.T) {
	c.GetCreds("creds", "../../")

	req, err := http.NewRequest("GET", "/v1/movies", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetMovies)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status. Got %v, Want %v", status, http.StatusOK)
	}
}

func TestDiscoverMovies(t *testing.T) {
	requests := []requestTest{
		{"/v1/discover", "Must pass a date", 400},
		{"/v1/discover?date=2021-05-06", `{"page":1,"results":[`, 200},
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
		if !strings.HasPrefix(rr.Body.String(), expectedBody) {
			t.Errorf("Handler returned wrong body. Got %v, Want %v", rr.Body.String()[0:16], expectedBody)
		}
	}
}

func TestGetTvMovies(t *testing.T) {
	requests := []requestTest{
		{"/v1/tv-movies", "Must pass a date", 400},
		{"/v1/tv-movies?date=2021-05-06", `[{"startTime":"2021-05-05T21:33Z",`, 200},
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
		if !strings.HasPrefix(rr.Body.String(), expectedBody) {
			t.Errorf("Handler returned wrong body. Got %v, Want %v", rr.Body.String()[0:16], expectedBody)
		}
	}
}

func TestGetTMSReq(t *testing.T) {
	// get creds
	// mock loc
	// make REAL request
	// test response
}
