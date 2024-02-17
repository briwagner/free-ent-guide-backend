package main

import (
	"fmt"
	"free-ent-guide-backend/models"
	"free-ent-guide-backend/models/modelstore"
	"free-ent-guide-backend/pkg/cred"
	"os"
	"slices"
	"strings"
)

var (
	c       cred.Cred
	DB      models.Store
	Querier *modelstore.Queries

	format = "2006-01-02"
)

func main() {
	commands := []string{"nhl", "mlb", "gup", "cache"}

	// Set-up application config.
	c.GetCreds("creds", ".")
	_, sqlc := models.Setup(c)
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

	switch cmd {
	case "nhl":
		handleNHL(Querier, os.Args)
	case "mlb":
		handleMLB(Querier, os.Args)
	case "cache":
		handleCache(Querier, os.Args)
	case "gup":
		handleGameUpdates(Querier, os.Args)
	}
}
