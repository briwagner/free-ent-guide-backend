package main

import (
	"errors"
	"fmt"
	"free-ent-guide-backend/models"
	"log"
	"time"
)

func handleMLB(tp TaskPayload, args []string) error {
	subCo := args[2]
	if subCo == "" {
		return errors.New("for MLB, must pass valid date or 'last'")
	}

	var ret string
	defer func() {
		if ret == "" {
			return
		}
		fmt.Println(ret)
		err := slackMessage(tp.Cred, ret)
		if err != nil {
			log.Println(err)
		}
	}()

	if subCo == "last" || subCo == "latest" {
		games, err := models.MLBGetLatestGames(tp.Querier)
		if err != nil {
			return fmt.Errorf("error fetching games %w", err)
		}

		if len(games) == 0 {
			log.Println("no games found")
			return nil
		}
		log.Printf("Got %d MLB games for %s\n", len(games), games[0].Gametime.Format(format))
		return nil
	}

	d, err := time.Parse(format, subCo)
	if err != nil {
		return fmt.Errorf("MLB game importer error: bad date for %s: %w", subCo, err)
	}
	ret, err = models.ImportMLB(tp.Querier, tp.client, d)
	if err != nil {
		return fmt.Errorf("MLB importer error for %s: %w", subCo, err)
	}

	return nil
}
