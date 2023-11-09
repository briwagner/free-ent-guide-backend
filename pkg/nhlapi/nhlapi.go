package nhlapi

import (
	"encoding/json"
	"fmt"
	"free-ent-guide-backend/models"
	"io"
	"log"
	"net/http"
	"time"

	"gorm.io/gorm/clause"
)

// GameSchedule is the payload for a single game.
type GameSchedule struct {
	ID       int       `json:"id"`
	GameType GameType  `json:"gameType"` // how does this translate?
	Gametime time.Time `json:"startTimeUTC"`
	Home     NHLTeam   `json:"homeTeam"`
	Away     NHLTeam   `json:"awayTeam"`
}

// GameDaySchedule is the playload for a single day, containing multiple games.
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
	ID   int    `json:"id"`
	Name string // any way to get this? is it needed?
	Link string // we don't seem to use this anyway
}

type GameType struct {
	Type string `json:"gameType"`
}

func (gt *GameType) UnmarshalJSON(b []byte) error {
	var val int
	if err := json.Unmarshal(b, &val); err != nil {
		return err
	}

	switch val {
	case 2:
		gt.Type = "Regular Season"
	case 99: // what are these other values??
		gt.Type = "Pre-season"
	case 100:
		gt.Type = "Playoff"
	default:
		gt.Type = "Game" // this should not happen
	}

	return nil
}

// ImportNHL is the 2023 go-native version to fetch NHL game schedule
// for a given week, and import games and teams to the DB.
func ImportNHL(store models.Store, startDate string) error {
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

	count := 0

	for _, days := range payload.Days {
		fmt.Printf("import %d NHL games for %s\n", days.NumberOfGames, days.Date)
		for _, g := range days.Games {

			game := models.NHLGame{
				GameID:      g.ID,
				Gametime:    g.Gametime,
				Description: g.GameType.Type,
			}

			// Map MLB team IDs to my database IDs.
			home := models.NHLTeam{TeamID: g.Home.ID}
			tx := store.FirstOrCreate(&home, models.NHLTeam{TeamID: g.Home.ID})
			if tx.Error != nil {
				log.Print(tx.Error)
				continue
			}
			game.HomeID = home.ID

			away := models.NHLTeam{TeamID: g.Away.ID}
			tx = store.FirstOrCreate(&away, models.NHLTeam{TeamID: g.Away.ID})
			if tx.Error != nil {
				log.Print(tx.Error)
				continue
			}
			game.VisitorID = away.ID

			tx = store.Omit(clause.Associations).Create(&game)
			if tx.Error != nil {
				log.Print(tx.Error)
				continue
			}

			// fmt.Printf("got game %v\n", game.GameID)
			// fmt.Printf("home ID %d | %d\n", game.HomeID, g.Home.ID)
			// fmt.Printf("away ID %d | %d\n", game.VisitorID, g.Away.ID)
			// fmt.Printf("gametime %v\n\n", game.Gametime)

			count++
		}
	}

	log.Printf("nhl import complete, adding %d games", count)

	return nil
}
