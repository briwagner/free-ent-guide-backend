package moviedb

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

func TestGetTrending(t *testing.T) {
	c.Setup()
	mdb := MovieDb{Key: c.Moviedb}
	mdb.GetTrending()

	if mdb.Status != http.StatusOK {
		t.Errorf("Fetch returned wrong status. Got %v, Wanted %v", mdb.Status, http.StatusOK)
	}
}

func TestGetDiscover(t *testing.T) {
	mdb := MovieDb{Key: c.Moviedb}
	dateParam := time.Now().Format("2006-01-02")
	mdb.GetDiscover(dateParam)

	if mdb.Status != http.StatusOK {
		t.Errorf("Fetch returned wrong status. Got %v, Wanted %v", mdb.Status, http.StatusOK)
	}
}

func TestGetToken(t *testing.T) {
	mdb := MovieDb{Key: c.Moviedb}
	mdb.GetToken()

	if mdb.Status != http.StatusOK {
		t.Errorf("Fetch returned wrong status. Got %v, Wanted %v", mdb.Status, http.StatusOK)
	}
	if mdb.Token.Success != true {
		t.Errorf("Token returned wrong status. Got %v, Wanted %v", mdb.Token.Success, true)
	}
}
