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

	var ret string
	defer func() {
		fmt.Println(ret)
		err := slackMessage(ret)
		if err != nil {
			log.Println(err)
		}
	}()

	if subCo == "last" {
		games, err := models.MLBGetLatestGames(q)
		if err != nil {
			log.Printf("error fetching games %s", err)
			return
		}

		if len(games) == 0 {
			log.Println("no games found")
			return
		}
		log.Printf("Got %d MLB games for %s\n", len(games), games[0].Gametime.Format(format))
		return
	}

	d, err := time.Parse(format, subCo)
	if err != nil {
		ret = fmt.Sprintf("MLB game importer error: bad date for %s: %s\n", subCo, err)
		return
	}
	ret, err = models.ImportMLB(q, d)
	if err != nil {
		ret = fmt.Sprintf("MLB importer error for %s: %s", subCo, err)
		return
	}
}
