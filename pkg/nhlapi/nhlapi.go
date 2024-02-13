package nhlapi

import (
	"encoding/json"
	"fmt"
	"free-ent-guide-backend/models"
	"free-ent-guide-backend/models/modelstore"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/go-sql-driver/mysql"
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

// Example to get stats for one team, using franchiseID.
// https://api.nhle.com/stats/rest/en/team/summary?isAggregate=false&cayenneExp=franchiseId=32%20and%20gameTypeId=2%20and%20seasonId%3C=20232024%20and%20seasonId%3E=20232024

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
	default:
		gt.Type = "Game" // this should not happen
	}

	return nil
}

// ImportNHL is the 2023 go-native version to fetch NHL game schedule
// for a given week, and import games and teams to the DB.
func ImportNHL(q *modelstore.Queries, startDate string) error {
	url := "https://api-web.nhle.com/v1/schedule/" + startDate

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	payload := &GameWeekPayload{}
	err = json.Unmarshal(data, payload)
	if err != nil {
		return err
	}

	var count, countErrs int

	for _, days := range payload.Days {
		fmt.Printf("import %d NHL games for %s\n", days.NumberOfGames, days.Date)
		daysGames := days.Games

		for _, g := range daysGames {
			var game models.NHLGame
			game.GameID = g.ID
			game.Gametime = g.Gametime
			game.Description = g.GameType.Type
			game.Status = g.GameState

			// Map MLB team IDs to my database IDs.
			home := &models.NHLTeam{}
			err := home.FindByTeamID(q, int64(g.Home.ID))
			// TODO we don't try to create missing teams, cuz there isn't a clear API endpoint to look them up.
			// We would want name, ID and tricode for that.
			if err != nil {
				countErrs++
				log.Printf("error finding home team %d for game %d: %s", g.Home.ID, g.ID, err)
				continue
			}
			game.HomeID = home.ID

			away := &models.NHLTeam{}
			err = away.FindByTeamID(q, int64(g.Away.ID))
			if err != nil {
				countErrs++
				log.Printf("error finding away team %d for game %d: %s", g.Away.ID, g.ID, err)
				continue
			}
			game.VisitorID = away.ID

			err = game.Create(q)
			if err != nil {
				countErrs++
				// Ignore duplicate entry in logs.
				if err.(*mysql.MySQLError).Number != 1062 {
					log.Printf("error creating game %d: %s", g.ID, err)
				}
				continue
			}

			// fmt.Printf("got game %v\n", game.GameID)
			// fmt.Printf("home ID %d | %d\n", game.HomeID, g.Home.ID)
			// fmt.Printf("away ID %d | %d\n", game.VisitorID, g.Away.ID)
			// fmt.Printf("gametime %v\n\n", game.Gametime)

			count++
		}
	}

	log.Printf("nhl import complete, adding %d games. Errors: %d", count, countErrs)

	return nil
}

type GetTeamsPayload struct {
	Data []NHLTeam `json:"data"`
}

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
