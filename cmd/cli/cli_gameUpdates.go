package main

import (
	"fmt"
	"free-ent-guide-backend/models"
	"free-ent-guide-backend/models/modelstore"
	"log"
	"strings"
)

func handleGameUpdates(q *modelstore.Queries, args []string) {
	date := args[2]

	var ret []string
	defer func() {
		retMsg := strings.Join(ret, " | ")
		err := slackMessage(retMsg)
		if err != nil {
			log.Println(err)
		}
	}()

	// Check MLB games first.
	mlbs := models.MLBGames{}
	err := mlbs.LoadByDate(q, date)
	if err != nil {
		log.Print(err)
		return
	}
	ret = append(ret, fmt.Sprintf("updating MLB games %d\n", len(mlbs)))
	log.Println(ret)

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
	ret2 := fmt.Sprintf("updating NHL games %d\n", len(nhls))
	ret = append(ret, ret2)
	log.Println(ret2)
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
