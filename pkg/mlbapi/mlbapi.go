package mlbapi

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

//go:embed divisions.json
var divisionsJSON []byte

const (
	baseURL    = "https://statsapi.mlb.com"
	amerLeague = "103"
	natLeague  = "104"
	season     = "2025"
)

// ImportDates fetches a single day schedule.
func ImportDates(client *http.Client, sd time.Time) (*MLBGameDay, error) {
	dt := sd.Format("2006-01-02")
	url := fmt.Sprintf("%s/api/v1/schedule?sportId=1&startDate=%s&endDate=%s", baseURL, dt, dt)
	return importDates(client, sd, url)
}

// Internal call to manage url for testing.
func importDates(client *http.Client, sd time.Time, url string) (*MLBGameDay, error) {
	gameday := &MLBGameDay{}

	var (
		tries int
		wait  time.Duration = time.Second * 5
		resp  *http.Response
		err   error
	)

	for {
		tries++
		resp, err = client.Get(url)
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) && tries < 3 { // max attempts
				time.Sleep(wait * 2) // double the time for each try
				continue
			}
			return gameday, err
		}
		break
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

func GetGameUpdate(client *http.Client, link string) (MLBGameUpdate, error) {
	var gameup MLBGameUpdate

	url := fmt.Sprintf("%s/%s", baseURL, link)
	resp, err := client.Get(url)
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
		return gameup, fmt.Errorf("error unmarshal MLB update for %s: %w", link, err)
	}

	return gameup, nil
}

func GetStandings(client *http.Client) (Standings, []error) {
	// testdata/mlb_standings.json // AL only

	var wg sync.WaitGroup
	var alStandings Standings
	var nlStandings Standings
	alURL := fmt.Sprintf("%s/%s?leagueId=%s&season=%s&standingsTypes=%s", baseURL, "api/v1/standings", amerLeague, season, "regularSeason")
	nlURL := fmt.Sprintf("%s/%s?leagueId=%s&season=%s&standingsTypes=%s", baseURL, "api/v1/standings", natLeague, season, "regularSeason")
	errors := make([]error, 2)

	// Do American League
	wg.Add(1)
	go func() {
		defer wg.Done()

		resp, err := client.Get(alURL)
		if err != nil {
			errors[0] = err
			return
		}

		defer resp.Body.Close()
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			errors[0] = err
			return
		}

		dec := json.NewDecoder(bytes.NewReader(data))
		dec.UseNumber()
		err = dec.Decode(&alStandings)
		if err != nil {
			errors[0] = err
			return
		}
	}()
	// Do National League
	wg.Add(1)
	go func() {
		defer wg.Done()

		resp, err := client.Get(nlURL)
		if err != nil {
			errors[1] = err
			return
		}

		defer resp.Body.Close()
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			errors[1] = err
			return
		}

		dec := json.NewDecoder(bytes.NewReader(data))
		dec.UseNumber()
		err = dec.Decode(&nlStandings)
		if err != nil {
			errors[1] = err
			return
		}
	}()

	wg.Wait()

	// Get both calls and just smash 'em together.
	alStandings.Records = append(alStandings.Records, nlStandings.Records...)

	if alStandings.RecordByTeam == nil {
		// TODO why does this initialize as nil?
		alStandings.RecordByTeam = make(map[string]*Record)
	}

	divisionsLU, divErr := GetDivisions()
	if divErr != nil {
		log.Print(divErr)
	}

	for _, rec := range alStandings.Records {
		for _, team := range rec.TeamRecords {
			if _, exists := alStandings.RecordByTeam[team.Team.Name]; exists {
				continue
			}
			if rec.Division.Name == "" && divErr == nil {
				div, found := divisionsLU[string(rec.Division.ID)]
				if found {
					rec.Division.Name = div.Name
				}
			}
			alStandings.RecordByTeam[team.Team.Name] = &rec
		}
	}

	return alStandings, errors
}

func GetTeamInfo(client *http.Client, id string) {
	// testdata/mlb_team.json

	// url := "https://statsapi.mlb.com/api/v1/teams/" + id
}

func GetDivisions() (map[string]Division, error) {
	// TODO check file and if fail, do actual http lookup.
	// url := "https://statsapi.mlb.com/api/v1/divisions?sportId=1"

	var divs DivisionPayload
	divsLU := make(map[string]Division)

	err := json.Unmarshal(divisionsJSON, &divs)
	if err != nil {
		return divsLU, err
	}

	for _, d := range divs.Divisions {
		divsLU[strconv.Itoa(d.ID)] = d
	}
	return divsLU, err
}
