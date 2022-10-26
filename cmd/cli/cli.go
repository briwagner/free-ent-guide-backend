package main

import (
	"fmt"
	"free-ent-guide-backend/models"
	"free-ent-guide-backend/pkg/cred"
	"free-ent-guide-backend/pkg/docker_importer"
	"log"
	"os"
)

var c cred.Cred

// var DB *gorm.DB
var DB models.Store

func main() {
	// Set-up application config.
	c.GetCreds("creds", ".")
	DB = models.Setup(c)

	if len(os.Args) <= 2 {
		fmt.Println("Must pass command: 'nhl' or 'mlb', along with a date e.g. 2022-11-21")
		return
	}

	cmd := os.Args[1]
	if cmd != "nhl" && cmd != "mlb" {
		fmt.Println("Command must be 'nhl' or 'mlb'")
		return
	}

	date := os.Args[2]

	switch cmd {
	case "nhl":
		err := docker_importer.ImportNHL(DB, date)
		if err != nil {
			log.Print(err)
			return
		}
	case "mlb":
		err := docker_importer.ImportMLB(DB, date)
		if err != nil {
			log.Print(err)
			return
		}
	}

}
