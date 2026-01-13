package main

import (
	"fmt"
	"io"
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
	url := fmt.Sprintf("/v1/movies?zip=%s", "60048")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	app := App{} // TODO needs setup

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.GetMovies)

	handler.ServeHTTP(rr, req)

	resp := rr.Result()
	sc := resp.StatusCode
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if status := sc; status != http.StatusOK {
		t.Errorf("Handler returned wrong status. Got %v, Want %v. Resp: %s", status, http.StatusOK, body)
	}
}

func TestDiscoverMovies(t *testing.T) {
	requests := []requestTest{
		{"/v1/discover", "Must pass a date", 400},
		{"/v1/discover?date=2021-05-06", `{"page":1,"results":[`, 200},
	}

	app := App{} // needs setup

	for _, tt := range requests {
		req, err := http.NewRequest("GET", tt.url, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.DiscoverMovies)

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

	app := App{} // TODO needs setup

	for _, tt := range requests {
		req, err := http.NewRequest("GET", tt.url, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.GetTvMovies)

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
