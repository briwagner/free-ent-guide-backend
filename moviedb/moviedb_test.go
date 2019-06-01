package moviedb

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"

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
	Key = c.Moviedb
}

func TestGetTrending(t *testing.T) {
	c.Setup()
	mdb := MovieDb{Key}
	s, _ := mdb.GetTrending()

	if s != http.StatusOK {
		t.Errorf("Fetch returned wrong status. Got %v, Wanted %v", s, http.StatusOK)
	}
}

func TestGetDiscover(t *testing.T) {
	mdb := MovieDb{Key}
	s, _ := mdb.GetDiscover()

	if s != http.StatusOK {
		t.Errorf("Fetch returned wrong status. Got %v, Wanted %v", s, http.StatusOK)
	}
}

func TestGetToken(t *testing.T) {
	mdb := MovieDb{Key}
	s, tok := mdb.GetToken()

	if s != http.StatusOK {
		t.Errorf("Fetch returned wrong status. Got %v, Wanted %v", s, http.StatusOK)
	}
	if tok.Success != true {
		t.Errorf("Token returned wrong status. Got %v, Wanted %v", tok.Success, true)
	}
}
