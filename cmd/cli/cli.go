package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"free-ent-guide-backend/models"
	"free-ent-guide-backend/models/modelstore"
	"free-ent-guide-backend/pkg/cred"
	"io"
	"net/http"
	"os"
	"slices"
	"strings"
)

var (
	c       cred.Cred
	Querier *modelstore.Queries

	format = "2006-01-02"
)

func main() {
	commands := []string{"nhl", "mlb", "gup", "cache", "discover"}

	// Set-up application config.
	c.GetCreds("creds", ".")
	sqlc := models.Setup(c)
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
	case "discover":
		handleDiscoverMovies(c, os.Args)
	}
}

type SlackBody struct {
	Text string `json:"text"`
}

func slackMessage(msg string) error {
	url := c.SlackURL

	var b strings.Builder
	_, err := b.WriteString(msg)
	if err != nil {
		return fmt.Errorf("error writing to slack %w", err)
	}

	var sb SlackBody
	sb.Text = b.String()
	data, err := json.Marshal(sb)
	if err != nil {
		return err
	}

	cli := &http.Client{}
	resp, err := cli.Post(url, "applicaton/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		by, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("slackbot error resp, failed to read body %d", resp.StatusCode)
		}
		return fmt.Errorf("slackbot error resp %d: %s", resp.StatusCode, string(by))
	}

	return nil
}
