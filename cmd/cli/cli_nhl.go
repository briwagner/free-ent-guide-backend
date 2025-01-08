package main

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"free-ent-guide-backend/models"
	"free-ent-guide-backend/models/modelstore"
	"free-ent-guide-backend/pkg/nhlapi"
	"log"
	"slices"
	"time"
)

func handleNHL(tp TaskPayload, args []string) error {
	subCo := args[2]
	if subCo == "" {
		return errors.New("for NHL, must pass valid date or 'last'")
	}

	var ret string
	defer func() {
		if ret == "" {
			return
		}
		// Raw slack message
		log.Printf("slack: %s", ret)
		err := slackMessage(tp.Cred, ret)
		if err != nil {
			log.Println(err)
		}
	}()

	if subCo == "last" {
		games, err := models.NHLGetLatestGames(tp.Querier)
		if err != nil {
			return fmt.Errorf("error getting last games: %w", err)
		}
		if len(games) == 0 {
			return errors.New("no games found")
		}
		log.Printf("Got %d NHL games on %s\n", len(games), games[0].Gametime.Format(format))
		return nil
	}

	// Use only as needed.
	if subCo == "teamseed" {
		err := seedNHLTeams(tp.Querier)
		if err != nil {
			log.Println(err)
		}
		return nil
	}

	// Importer
	d, err := time.Parse(format, subCo)
	if err != nil {
		return fmt.Errorf("NHL game importer error: bad date for '%s'. Did you mean 'last'?", subCo)
	}
	ret, err = models.ImportNHL(tp.Querier, d.Format(format))
	if err != nil {
		return fmt.Errorf("NHL importer error for %s: %w", subCo, err)
	}

	return nil
}

//go:embed nhl_teams.json
var nhlTeamData []byte

func seedNHLTeams(q *modelstore.Queries) error {
	// These were taken from old api, based on teams from the 2023-24 game schedule.
	activeIDs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 28, 29, 30, 52, 53, 54, 55, 59, 87, 88, 89, 90, 99}

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

			err := dbTeam.Create(q)
			if err != nil {
				return err
			}
			counter++
		}
	}

	log.Printf("added %d teams", counter)
	return nil
}
