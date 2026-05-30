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
	"time"
)

//go:embed divisions.json
var divisionsJSON []byte // TODO update this every year

const (
	baseURL    = "http://statsapi.mlb.com" // stopped using https bc 406 error, assume rate-limit
	amerLeague = "103"
	natLeague  = "104"
	season     = "2026" // TODO update this every year!!
)

// ImportDates fetches a single day schedule.
func ImportDates(ctx context.Context, client *http.Client, sd time.Time) (*MLBGameDay, error) {
	dt := sd.Format("2006-01-02")
	url := fmt.Sprintf("%s/api/v1/schedule?sportId=1&startDate=%s&endDate=%s", baseURL, dt, dt)
	return importDates(ctx, client, sd, url)
}

// call to statsapi
func importDates(ctx context.Context, client *http.Client, sd time.Time, url string) (*MLBGameDay, error) {
	gameday := &MLBGameDay{}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return gameday, err
	}

	// TODO added these bc of rate-limit above
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:146.0) Gecko/20100101 Firefox/146.0")

	resp, err := client.Do(req)
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

func GetGameUpdate(client *http.Client, link string) (MLBGameUpdate, error) {
	var gameup MLBGameUpdate

	url := fmt.Sprintf("%s/%s", baseURL, link)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return gameup, err
	}

	resp, err := client.Do(req)
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

func GetStandings(ctx context.Context, client *http.Client) (Standings, error) {
	// testdata/mlb_standings.json // AL only

	var (
		alStandings, nlStandings Standings
	)
	alURL := fmt.Sprintf("%s/%s?leagueId=%s&season=%s&standingsTypes=%s", baseURL, "api/v1/standings", amerLeague, season, "regularSeason")
	nlURL := fmt.Sprintf("%s/%s?leagueId=%s&season=%s&standingsTypes=%s", baseURL, "api/v1/standings", natLeague, season, "regularSeason")
	errs := make([]error, 2)

	{
		// Do American League
		req, err := http.NewRequestWithContext(ctx, "GET", alURL, nil)
		if err != nil {
			errs[0] = err
		}
		resp, err := client.Do(req)
		if err != nil {
			errs[0] = err
		}

		defer resp.Body.Close()
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			errs[0] = err
		}

		dec := json.NewDecoder(bytes.NewReader(data))
		dec.UseNumber()
		err = dec.Decode(&alStandings)
		if err != nil {
			errs[0] = err
		}
	}

	{
		// Do National League
		req, err := http.NewRequestWithContext(ctx, "GET", nlURL, nil)
		if err != nil {
			errs[0] = err
		}
		resp, err := client.Do(req)
		if err != nil {
			errs[1] = err
		}

		defer resp.Body.Close()
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			errs[1] = err
		}

		dec := json.NewDecoder(bytes.NewReader(data))
		dec.UseNumber()
		err = dec.Decode(&nlStandings)
		if err != nil {
			errs[1] = err
		}
	}

	// Get both calls and just smash 'em together.
	alStandings.Records = append(alStandings.Records, nlStandings.Records...)

	divisionsLUT, divErr := GetDivisions()
	if divErr != nil {
		// what is this error and how can we continue with it?
		log.Print(divErr)
		errs = append(errs, divErr)
	}
	if len(divisionsLUT) == 0 {
		errs = append(errs, errors.New("no divisions found in LUT"))
		return alStandings, errors.Join(errs...)
	}

	alStandings.Finalize(divisionsLUT, divErr)

	return alStandings, errors.Join(errs...)
}

func GetTeamInfo(client *http.Client, id string) {
	// testdata/mlb_team.json

	// url := "https://statsapi.mlb.com/api/v1/teams/" + id
}

// GetDivisions returns a LUT to use in grouping teams for standings.
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
