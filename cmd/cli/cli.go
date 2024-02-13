package main

import (
	"encoding/json"
	"fmt"
	"free-ent-guide-backend/models"
	"free-ent-guide-backend/models/modelstore"
	"free-ent-guide-backend/pkg/cred"
	"free-ent-guide-backend/pkg/docker_importer"
	"free-ent-guide-backend/pkg/nhlapi"
	"log"
	"os"
	"slices"
	"strings"
	"time"
)

var (
	c       cred.Cred
	DB      models.Store
	Querier *modelstore.Queries
)

func main() {
	commands := []string{"nhl", "mlb", "gup", "cache"}

	// Set-up application config.
	c.GetCreds("creds", ".")
	DB, sqlc := models.Setup(c)
	Querier = modelstore.New(sqlc)

	if len(os.Args) <= 2 {
		fmt.Printf(`
Available commands: %s.
'nhl', 'mlb' or 'gup' along with a date e.g. 2022-11-21.
Or 'nhl', 'mlb' with 'last'.
Or 'cache' with one of: show, stale, clear, wipe.

`, strings.Join(commands, ", "))
		return
	}

	cmd := os.Args[1]
	if !slices.Contains(commands, cmd) {
		fmt.Printf("Command must be one of %s\n", strings.Join(commands, ", "))
		return
	}

	// Date format.
	format := "2006-01-02"

	switch cmd {
	case "nhl":
		subCo := os.Args[2]
		if subCo == "" {
			fmt.Println("For NHL, must pass valid date or 'last'.")
			return
		}

		if subCo == "last" {
			games, err := models.GetLatestGames(Querier)
			if err != nil {
				fmt.Printf("Error getting last games: %s\n", err)
				return
			}
			fmt.Printf("Got %d NHL games on %s\n", len(games), games[0].Gametime.Format(format))
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
		err = nhlapi.ImportNHL(Querier, d.Format(format))
		if err != nil {
			log.Print(err)
			return
		}
	case "mlb":
		date := os.Args[2]
		err := docker_importer.ImportMLB(DB, date)
		if err != nil {
			log.Print(err)
			return
		}
	case "cache":
		op := os.Args[2]
		ops := []string{"show", "stale", "clear", "wipe"}
		if !slices.Contains(ops, op) {
			fmt.Printf("For 'cache', must pass one of %s\n", strings.Join(ops, ", "))
		}
		cs := &models.Caches{}

		if op == "show" {
			err := cs.ShowAll(Querier)
			if err != nil {
				log.Print(err)
				return
			}
			if len(*cs) == 0 {
				log.Print("no records found")
				return
			}
			for _, c := range *cs {
				fmt.Println(c)
			}
			return
		}

		if op == "stale" {
			err := cs.ShowStale(time.Now(), Querier)
			if err != nil {
				log.Print(err)
				return
			}
			if len(*cs) == 0 {
				log.Print("no records found")
				return
			}
			for _, c := range *cs {
				fmt.Println(c)
			}
			return
		}

		if op == "clear" {
			rows, err := cs.DropStale(time.Now(), Querier)
			if err != nil {
				log.Print(err)
				return
			}
			log.Printf("Dropped %d records.\n", rows)
		}

		if op == "wipe" {
			rows, err := cs.DropAll(Querier)
			if err != nil {
				log.Print(err)
				return
			}
			log.Printf("%d caches wiped\n", rows)
		}

	case "gup":
		date := os.Args[2]

		// Check MLB games first.
		// mlbs := models.MLBGames{}
		// err := mlbs.LoadByDate(date, DB)
		// if err != nil {
		// 	log.Print(err)
		// 	return
		// }
		// log.Printf("updating MLB games %d\n", len(mlbs))
		// for _, mlb := range mlbs {
		// 	err := mlb.UpdateScore(DB)
		// 	if err != nil {
		// 		log.Println(err)
		// 	}
		// }

		// Do same for NHL.
		nhls := models.NHLGames{}
		err := nhls.LoadByDate(Querier, date)
		if err != nil {
			log.Print(err)
			return
		}
		log.Printf("updating NHL games %d\n", len(nhls))
		for _, nhl := range nhls {
			if nhl.Status == "Final" {
				continue
			}
			err := nhl.UpdateScore(Querier)
			if err != nil {
				log.Println(err)
			}
		}
	}

}

func seedNHLTeams() error {
	// These were taken from old api, based on teams from the 2023-24 game schedule.
	activeIDs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 28, 29, 30, 52, 53, 54, 55, 87, 88, 89, 90, 99}

	data, err := os.ReadFile("pkg/nhlapi/testdata/nhl_teams.json")
	if err != nil {
		return err
	}

	var payload nhlapi.GetTeamsPayload
	err = json.Unmarshal(data, &payload)
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
