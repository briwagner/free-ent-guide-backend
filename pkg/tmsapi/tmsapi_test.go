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
var Key string

type Cred struct {
	Tms     string
	Moviedb string
}

func (c *Cred) Setup() {
	yamlFile, err := ioutil.ReadFile("../creds.yaml")
	if err != nil {
		log.Fatalf("File not loading: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("No credentials %v", err)
	}
	Key = c.Tms
}

func TestGetCinema(t *testing.T) {
	c.Setup()
	tms := TmsApi{Key}
	d := time.Now().Format("2006-01-02")
	s, _ := tms.GetCinema("60048", d)

	if s != http.StatusOK {
		t.Errorf("Fetch returned wrong status. Got %v, Wanted %v", s, http.StatusOK)
	}
}

func TestGetTvMovies(t *testing.T) {
	c.Setup()
	tms := TmsApi{Key}
	d := time.Now().Format("2006-01-02")
	s, _ := tms.GetTvMovies(d)

	if s != http.StatusOK {
		t.Errorf("Fetch returned wrong status. Got %v, Wanted %v", s, http.StatusOK)
	}
}

func TestGetTvSports(t *testing.T) {
	c.Setup()
	tms := TmsApi{Key}
	d := time.Now().Format("2006-01-02")
	s, _ := tms.GetTvSports(d)

	if s != http.StatusOK {
		t.Errorf("Fetch returned wrong status. Got %v, Wanted %v", s, http.StatusOK)
	}
}
