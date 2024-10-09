package nhlapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// GameSchedule is the payload for a single game.
type GameSchedule struct {
	ID        int       `json:"id"`
	GameType  GameType  `json:"gameType"` // how does this translate?
	Gametime  time.Time `json:"startTimeUTC"`
	GameState string    `json:"gameState"`
	Home      NHLTeam   `json:"homeTeam"`
	Away      NHLTeam   `json:"awayTeam"`
}

// GameDaySchedule is the payload for a single day, containing multiple games.
type GameDaySchedule struct {
	Date          string         `json:"date"`
	NumberOfGames int            `json:"numberOfGames"`
	Games         []GameSchedule `json:"games"`
}

// GameWeekPayload is the payload for a week.
type GameWeekPayload struct {
	Days []GameDaySchedule `json:"gameWeek"`
}

type NHLTeam struct {
	ID          int    `json:"id"`          // this is the one mapped to Games
	FranchiseID int    `json:"franchiseId"` // this can be used for querying on summary, see below.
	Tricode     string `json:"triCode"`
	Name        string `json:"fullName"`
	Link        string // we don't seem to use this anyway
}

// TODO is this v1 only?
// Example to get stats for one team, using franchiseID.
// https://api.nhle.com/stats/rest/en/team/summary?isAggregate=false&cayenneExp=franchiseId=32%20and%20gameTypeId=2%20and%20seasonId%3C=20232024%20and%20seasonId%3E=20232024

// List of Franchises, V2
// https://api.nhle.com/stats/rest/en/franchise

type GameType struct {
	Type string `json:"gameType"`
}

func (gt *GameType) UnmarshalJSON(b []byte) error {
	var val int
	if err := json.Unmarshal(b, &val); err != nil {
		return err
	}

	switch val {
	case 1: //
		gt.Type = "Pre-season"
	case 2:
		gt.Type = "Regular Season"
	case 3: // what are these other values??
		gt.Type = "Playoff"
	case 4:
		gt.Type = "All-Star Game"
	default:
		gt.Type = "Game" // this should not happen
	}

	return nil
}

// ImportNHL is the 2023 go-native version to fetch NHL game schedule
// for a given week, and import games and teams to the DB.
func ImportNHL(startDate string) (*GameWeekPayload, error) {
	url := "https://api-web.nhle.com/v1/schedule/" + startDate
	payload := &GameWeekPayload{}

	resp, err := http.Get(url)
	if err != nil {
		return payload, err
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return payload, err
	}

	err = json.Unmarshal(data, payload)
	if err != nil {
		return payload, err
	}

	return payload, nil
}

type GetTeamsPayload struct {
	Data []NHLTeam `json:"data"`
}

type (
	GetTeamPayload struct {
		Teams []GetTeamTeam `json:"teams"`
	}

	GetTeamTeam struct {
		Tricode string `json:"tricode"`
		TeamID  int    `json:"teamId"`
		Name    struct {
			Default string `json:"default"`
		} `json:"name"`
	}
)

// Assume we fetch a single team and convert it.
func (gt *GetTeamPayload) ToTeam() NHLTeam {
	nhlt := NHLTeam{}

	if len(gt.Teams) == 0 {
		return nhlt
	}

	nhlt.Tricode = gt.Teams[0].Tricode
	nhlt.ID = gt.Teams[0].TeamID
	nhlt.Name = gt.Teams[0].Name.Default

	return nhlt
}

func GetTeam(id string) (NHLTeam, error) {
	url := "https://api-web.nhle.com/v1/meta?teams=" + id

	t := NHLTeam{}
	resp, err := http.Get(url)
	if err != nil {
		return t, fmt.Errorf("failed to get team %s: %s", id, err)
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return t, fmt.Errorf("error reading response for nhl team: %s", err)
	}

	pl := GetTeamPayload{}
	err = json.Unmarshal(data, &pl)
	if err != nil {
		return t, fmt.Errorf("error unmarshalling nhl team: %s", err)
	}

	team := pl.ToTeam()

	return team, nil
}

// TODO unused. Convert this to v2 api, for team ID or IDs, e.g. https://api-web.nhle.com/v1/meta?teams=59,14,26,28
func GetTeams() ([]NHLTeam, error) {
	var teams []NHLTeam

	url := "https://api.nhle.com/stats/rest/en/team/summary"
	res, err := http.Get(url)
	if err != nil {
		log.Printf("error fetching NHL teams: %s", err)
		return teams, err
	}

	defer res.Body.Close()
	var payload GetTeamsPayload
	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("error reading NHL teams response: %s", err)
		return teams, err
	}

	err = json.Unmarshal(data, &payload)
	if err != nil {
		log.Printf("error unmarshal NHL teams: %s", err)
		return teams, err
	}

	return teams, nil
}

// GetUpdate makes an api request to get game update to
// merge with scheduled game info from the database.
func GetUpdate(gameID int) (NHLGameUpdate, error) {
	var gu NHLGameUpdate

	base := "https://api-web.nhle.com/v1/gamecenter"
	url := fmt.Sprintf("%s/%d/play-by-play", base, gameID)
	resp, err := http.Get(url)
	if err != nil {
		return gu, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return gu, err
	}

	err = json.Unmarshal(body, &gu)
	if err != nil {
		return gu, err
	}

	return gu, nil
}

// NHLGameUpdate wraps info on a scheduled game, live status (if applicable)
// and scores.
type NHLGameUpdate struct {
	ID           int64
	Status       string
	Period       int64
	VisitorScore int64
	HomeScore    int64
}

// As of 11-2023, NHL api has multiple versions of 'Final'.
// For complete games, it returns "OFF".
// We set a custom value of 'Final' and send as signal to front-end.
func setGameState(st string) string {
	switch st {
	case "OFF", "OVER", "FINAL":
		return "Final"
	}
	return st
}

// UnmarshalJSON is required to extract the minimal data needed on the frontend.
func (g *NHLGameUpdate) UnmarshalJSON(b []byte) error {
	var cg map[string]interface{}

	// Get ID.
	dec := json.NewDecoder(bytes.NewReader(b))
	dec.UseNumber()
	err := dec.Decode(&cg)
	if err != nil {
		return fmt.Errorf("error decoding: %w", err)
	}
	id, err := cg["id"].(json.Number).Int64()
	if err != nil {
		return fmt.Errorf("error parsing ID: %w", err)
	}
	g.ID = id

	st := cg["gameState"].(string)
	g.Status = setGameState(st)

	// Only check for live or past games, else this key is not found.
	if g.Status == "LIVE" || g.Status == "Final" {
		periodDesc := cg["periodDescriptor"].(map[string]interface{})
		per, err := periodDesc["number"].(json.Number).Int64()
		if err == nil {
			g.Period = per
		}
	}

	// Set scores.
	home := cg["homeTeam"].(map[string]interface{})
	hScore, err := home["score"].(json.Number).Int64()
	if err != nil {
		return err
	}
	g.HomeScore = hScore

	away := cg["awayTeam"].(map[string]interface{})
	aScore, err := away["score"].(json.Number).Int64()
	if err != nil {
		return err
	}
	g.VisitorScore = aScore

	return nil
}
