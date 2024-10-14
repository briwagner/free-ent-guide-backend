package mlbapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"time"
)

const baseURL = "https://statsapi.mlb.com/"

type MLBGameDay struct {
	Dates []struct {
		Date  string    `json:"date"`
		Games []MLBGame `json:"games"`
	} `json:"dates"`
}

type MLBGame struct {
	Time        time.Time    `json:"gameDate"`
	Link        string       `json:"link"`
	GameID      int          `json:"gamePk"`
	Description string       `json:"seriesDescription"` // 'Regular Season', etc.
	Status      StatusString `json:"status"`
	Teams       MLBTeams     `json:"teams"`
}

type StatusString struct {
	Status string
}

func (ss *StatusString) UnmarshalJSON(data []byte) error {
	var vals map[string]interface{}
	if err := json.Unmarshal(data, &vals); err != nil {
		return err
	}
	state, ok := vals["detailedState"]
	if ok {
		ss.Status = state.(string)
	}

	return nil
}

type MLBTeams struct {
	Home MLBTeam `json:"home"`
	Away MLBTeam `json:"away"`
}

type MLBTeam struct {
	Score int    `json:"-"`
	Name  string `json:"name"`
	ID    int    `json:"id"`
	Link  string `json:"link"`
}

func (t *MLBTeam) UnmarshalJSON(data []byte) error {
	var vals map[string]interface{}
	if err := json.Unmarshal(data, &vals); err != nil {
		return err
	}

	if sc, ok := vals["score"]; ok {
		t.Score = int(math.Round(sc.(float64)))
	}

	team := vals["team"].(map[string]interface{})
	if n, ok := team["name"]; ok {
		t.Name = n.(string)
	}
	if l, ok := team["link"]; ok {
		t.Link = l.(string)
	}
	if id, ok := team["id"]; ok {
		t.ID = int(math.Round(id.(float64)))
	}

	return nil
}

// ImportDates fetches a single day schedule.
func ImportDates(sd time.Time) (*MLBGameDay, error) {
	gameday := &MLBGameDay{}

	dt := sd.Format("2006-01-02")
	url := fmt.Sprintf("%s/api/v1/schedule?sportId=1&startDate=%s&endDate=%s", baseURL, dt, dt)
	resp, err := http.Get(url)
	if err != nil {
		return gameday, err
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return gameday, err
	}

	err = json.Unmarshal(data, gameday)
	if err != nil {
		return gameday, err
	}

	return gameday, nil
}

type MLBGameUpdate struct {
	GamePK       int64
	Game         MLBGame
	Status       string
	Inning       int
	HomeScore    int
	VisitorScore int
}

func (gu *MLBGameUpdate) UnmarshalJSON(data []byte) error {
	var vals map[string]interface{}
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	err := dec.Decode(&vals)
	if err != nil {
		return err
	}

	id, err := vals["gamePk"].(json.Number).Int64()
	if err != nil {
		return err
	}
	gu.GamePK = id

	gd := vals["gameData"].(map[string]interface{})
	st := gd["status"].(map[string]interface{})
	// TODO add nil check
	gu.Status = st["detailedState"].(string)

	ld := vals["liveData"].(map[string]interface{})
	ls := ld["linescore"].(map[string]interface{})
	// If not live, the currentInning prop is not found.
	if _, ok := ls["currentInning"]; ok {
		inning, err := ls["currentInning"].(json.Number).Int64()
		if err != nil {
			return err
		}
		gu.Inning = int(inning)
	}

	ts := ls["teams"].(map[string]interface{})
	ht := ts["home"].(map[string]interface{})
	if _, ok := ht["runs"]; ok {
		hg, err := ht["runs"].(json.Number).Int64()
		if err != nil {
			return err
		}
		gu.HomeScore = int(hg)
	}

	vt := ts["away"].(map[string]interface{})
	if _, ok := ht["runs"]; ok {
		vg, err := vt["runs"].(json.Number).Int64()
		if err != nil {
			return err
		}
		gu.VisitorScore = int(vg)
	}
	return nil
}

func GetGameUpdate(link string) (MLBGameUpdate, error) {
	var gameup MLBGameUpdate

	url := fmt.Sprintf("%s/%s", baseURL, link)
	resp, err := http.Get(url)
	if err != nil {
		return gameup, fmt.Errorf("error fetching MLB update: %w", err)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return gameup, fmt.Errorf("error reading MLB update: %w", err)
	}

	err = json.Unmarshal(data, &gameup)
	if err != nil {
		return gameup, fmt.Errorf("error unmarshal MLB update: %w", err)
	}

	return gameup, nil
}
