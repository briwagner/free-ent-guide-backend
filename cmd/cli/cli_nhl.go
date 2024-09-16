package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"free-ent-guide-backend/models"
	"free-ent-guide-backend/models/modelstore"
	"free-ent-guide-backend/pkg/nhlapi"
	"log"
	"slices"
	"time"
)

func handleNHL(q *modelstore.Queries, args []string) {
	subCo := args[2]
	if subCo == "" {
		fmt.Println("For NHL, must pass valid date or 'last'.")
		return
	}

	var ret string
	defer func() {
		log.Println(ret)
		err := slackMessage(ret)
		if err != nil {
			log.Println(err)
		}
	}()

	if subCo == "last" {
		games, err := models.NHLGetLatestGames(q)
		if err != nil {
			ret = fmt.Sprintf("Error getting last games: %s\n", err)
			return
		}
		if len(games) == 0 {
			fmt.Println("no games found")
			return
		}
		ret = fmt.Sprintf("Got %d NHL games on %s\n", len(games), games[0].Gametime.Format(format))
		return
	}

	// Use only as needed.
	if subCo == "teamseed" {
		err := seedNHLTeams()
		if err != nil {
			log.Println(err)
		}
		return
	}

	// Importer
	d, err := time.Parse(format, subCo)
	if err != nil {
		fmt.Printf("NHL game importer error: bad date for '%s'. Did you mean 'last'?\n", subCo)
		return
	}
	err = models.ImportNHL(q, d.Format(format))
	if err != nil {
		log.Print(err)
		return
	}
}

//go:embed nhl_teams.json
var nhlTeamData []byte

func seedNHLTeams() error {
	// These were taken from old api, based on teams from the 2023-24 game schedule.
	activeIDs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 28, 29, 30, 52, 53, 54, 55, 87, 88, 89, 90, 99}

	data := nhlTeamData
	// data, err := os.ReadFile("pkg/nhlapi/testdata/nhl_teams.json")
	// if err != nil {
	// 	return err
	// }

	var payload nhlapi.GetTeamsPayload
	err := json.Unmarshal(data, &payload)
	if err != nil {
		return err
	}

	counter := 0
	for _, apiTeam := range payload.Data {
		if slices.Contains(activeIDs, apiTeam.ID) {
			var dbTeam models.NHLTeam
			dbTeam.Name = apiTeam.Name
			dbTeam.TeamID = apiTeam.ID
			dbTeam.Tricode = apiTeam.Tricode
			dbTeam.FranchiseID = apiTeam.FranchiseID
			dbTeam.UpdatedAt = time.Now()

			err := dbTeam.Create(Querier)
			if err != nil {
				return err
			}
			counter++
		}
	}

	log.Printf("added %d teams", counter)
	return nil
}
