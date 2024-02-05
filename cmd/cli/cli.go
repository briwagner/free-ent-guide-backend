package main

import (
	"fmt"
	"free-ent-guide-backend/models"
	"free-ent-guide-backend/pkg/cred"
	"free-ent-guide-backend/pkg/docker_importer"
	"free-ent-guide-backend/pkg/nhlapi"
	"log"
	"os"
	"slices"
	"strings"
	"time"
)

var c cred.Cred

var DB models.Store

func main() {
	commands := []string{"nhl", "mlb", "gup", "cache"}

	// Set-up application config.
	c.GetCreds("creds", ".")
	DB = models.Setup(c)

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
			games, err := models.GetLatestGames(DB)
			if err != nil {
				fmt.Printf("Error getting last games: %s\n", err)
				return
			}
			if len(games) == 0 {
				fmt.Println("no games found")
				return
			}
			fmt.Printf("Got %d NHL games on %s\n", len(games), games[0].Gametime.Format(format))
			return
		}

		// Importer
		d, err := time.Parse(format, subCo)
		if err != nil {
			fmt.Printf("NHL game importer error: bad date for '%s'. Did you mean 'last'?\n", subCo)
			return
		}
		err = nhlapi.ImportNHL(DB, d.Format(format))
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
			if nhl.Status == "Final" {
				continue
			}
			err := nhl.UpdateScore(DB)
			if err != nil {
				log.Println(err)
			}
		}
	}

}
