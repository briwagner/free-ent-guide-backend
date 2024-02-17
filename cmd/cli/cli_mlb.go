package main

import (
	"fmt"
	"free-ent-guide-backend/models"
	"free-ent-guide-backend/models/modelstore"
	"log"
	"time"
)

func handleMLB(q *modelstore.Queries, args []string) {
	subCo := args[2]
	if subCo == "" {
		fmt.Println("For MLB, must pass valid date or 'last'.")
		return
	}

	if subCo == "last" {
		games, err := models.MLBGetLatestGames(q)
		if err != nil {
			fmt.Printf("error fetching games %s", err)
		}

		if len(games) == 0 {
			fmt.Println("no games found")
			return
		}
		fmt.Printf("Got %d MLB games for %s\n", len(games), games[0].Gametime.Format(format))
		return
	}

	d, err := time.Parse(format, subCo)
	if err != nil {
		log.Printf("MLB game importer error: bad date for %s: %s\n", subCo, err)
		return
	}
	err = models.ImportMLB(q, d)
	if err != nil {
		log.Printf("MLB importer error: %s\n", err)
		return
	}
}
