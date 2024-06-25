package main

import (
	"free-ent-guide-backend/models"
	"free-ent-guide-backend/models/modelstore"
	"log"
)

func handleGameUpdates(q *modelstore.Queries, args []string) {
	date := args[2]

	// Check MLB games first.
	mlbs := models.MLBGames{}
	err := mlbs.LoadByDate(q, date)
	if err != nil {
		log.Print(err)
		return
	}
	log.Printf("updating MLB games %d\n", len(mlbs))
	for _, mlb := range mlbs {
		err := mlb.UpdateScore(q)
		if err != nil {
			log.Println(err)
		}
	}

	// Do same for NHL.
	nhls := models.NHLGames{}
	err = nhls.LoadByDate(q, date)
	if err != nil {
		log.Print(err)
		return
	}
	log.Printf("updating NHL games %d\n", len(nhls))
	for _, nhl := range nhls {
		if nhl.Status == "Final" {
			continue
		}
		err := nhl.UpdateScore(q)
		if err != nil {
			log.Println(err)
		}
	}
}
