package tmsapi

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"

	"gopkg.in/yaml.v2"
)

var c Cred

type Cred struct {
	Tms     string
	Moviedb string
}

func (c *Cred) Setup() {
	yamlFile, err := ioutil.ReadFile("../../creds.yaml")
	if err != nil {
		log.Fatalf("File not loading: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("No credentials %v", err)
	}
}

func TestGetCinema(t *testing.T) {
	c.Setup()
	tms := TmsApi{Key: c.Tms}
	d := time.Now().Format("2006-01-02")
	tms.GetCinema("60048", d)

	if tms.Status != http.StatusOK {
		t.Errorf("Fetch returned wrong status. Got %v, Wanted %v", tms.Status, http.StatusOK)
	}
}

func TestGetTvMovies(t *testing.T) {
	// c.Setup()
	tms := TmsApi{Key: c.Tms}
	d := time.Now().Format("2006-01-02")
	tms.GetTvMovies(d, lineup)

	if tms.Status != http.StatusOK {
		t.Errorf("Fetch returned wrong status. Got %v, Wanted %v", tms.Status, http.StatusOK)
	}
}

func TestGetTvSports(t *testing.T) {
	// c.Setup()
	tms := TmsApi{Key: c.Tms}
	d := time.Now().Format("2006-01-02")
	tms.GetTvSports(d, lineup)

	if tms.Status != http.StatusOK {
		t.Errorf("Fetch returned wrong status. Got %v, Wanted %v", tms.Status, http.StatusOK)
	}
}
