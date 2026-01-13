package main

import (
	"context"
	"errors"
	"fmt"
	"free-ent-guide-backend/models"
	"log"
	"log/slog"
	"time"
)

func handleMLB(ctx context.Context, l *slog.Logger, tp TaskPayload, args []string) error {
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
			l.Error("error handleMlb", "error", err)
		}
	}()

	if subCo == "last" || subCo == "latest" {
		games, err := models.MLBGetLatestGames(ctx, tp.Querier)
		if err != nil {
			return fmt.Errorf("error fetching games %w", err)
		}

		if len(games) == 0 {
			l.Warn("no mlb games found", "date", subCo)
			return nil
		}
		l.Info("got MLB games", "amount", len(games), "date", games[0].Gametime.Format(format))
		return nil
	}

	d, err := time.Parse(format, subCo)
	if err != nil {
		return fmt.Errorf("MLB game importer error: bad date for %s: %w", subCo, err)
	}
	ret, err = models.ImportMLB(ctx, tp.Querier, tp.client, d)
	if err != nil {
		return fmt.Errorf("MLB importer error for %s: %w", subCo, err)
	}

	return nil
}
