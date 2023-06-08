package main

import (
	"fmt"
	"free-ent-guide-backend/models"
	"free-ent-guide-backend/pkg/cred"
	"free-ent-guide-backend/pkg/docker_importer"
	"log"
	"os"
	"time"
)

var c cred.Cred

// var DB *gorm.DB
var DB models.Store

func main() {
	// Set-up application config.
	c.GetCreds("creds", ".")
	DB = models.Setup(c)

	if len(os.Args) <= 2 {
		fmt.Println("Must pass command: 'nhl', 'mlb' or 'gup', along with a date e.g. 2022-11-21. \nOr 'cache' with one of: show, stale, clear, wipe.")
		return
	}

	cmd := os.Args[1]
	if cmd != "nhl" && cmd != "mlb" && cmd != "cache" && cmd != "gup" {
		fmt.Println("Command must be 'nhl' or 'mlb' or 'cache'")
		return
	}

	switch cmd {
	case "nhl":
		date := os.Args[2]
		err := docker_importer.ImportNHL(DB, date)
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
		cs := models.Caches{}

		if op == "show" {
			err := cs.ShowAll(DB)
			if err != nil {
				log.Print(err)
				return
			}
			if len(cs) == 0 {
				log.Print("no records found")
				return
			}
			for _, c := range cs {
				fmt.Println(c)
			}
			return
		}

		if op == "stale" {
			err := cs.ShowStale(time.Now(), DB)
			if err != nil {
				log.Print(err)
				return
			}
			if len(cs) == 0 {
				log.Print("no records found")
				return
			}
			for _, c := range cs {
				fmt.Println(c)
			}
			return
		}

		if op == "clear" {
			err := cs.ShowStale(time.Now(), DB)
			if err != nil {
				log.Print(err)
				return
			}

			count := len(cs)
			if count == 0 {
				log.Println("no stale records found")
				return
			}
			err = cs.DropStale(DB)
			if err != nil {
				log.Print(err)
				return
			}
			log.Printf("Dropped %d records.\n", count)
		}

		if op == "wipe" {
			err := cs.DropAll(DB)
			if err != nil {
				log.Print(err)
				return
			}
			log.Println("caches wiped")
		}

	case "gup":
		date := os.Args[2]

		// Check MLB games first.
		mlbs := models.MLBGames{}
		err := mlbs.LoadByDate(date, DB)
		if err != nil {
			log.Print(err)
			return
		}
		log.Printf("updating MLB games %d\n", len(mlbs))
		for _, mlb := range mlbs {
			err := mlb.UpdateScore(DB)
			if err != nil {
				log.Println(err)
			}
		}

		// Do same for NHL.
		nhls := models.NHLGames{}
		err = nhls.LoadByDate(date, DB)
		if err != nil {
			log.Print(err)
			return
		}
		log.Printf("updating NHL games %d\n", len(nhls))
		for _, nhl := range nhls {
			err := nhl.UpdateScore(DB)
			if err != nil {
				log.Println(err)
			}
		}
	}

}
