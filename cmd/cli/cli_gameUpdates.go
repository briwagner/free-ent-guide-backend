package main

import (
	"context"
	"errors"
	"fmt"
	"free-ent-guide-backend/models"
	"log"
	"log/slog"
)

func handleGameUpdates(ctx context.Context, l *slog.Logger, tp TaskPayload, args []string) error {
	date := args[2]

	var ret string
	defer func() {
		if ret == "" {
			return
		}
		err := slackMessage(tp.Cred, ret)
		if err != nil {
			log.Println(err)
		}
	}()

	// TODO consider passing leagues in here to skip MLB during winter, for example.

	// Check MLB games first.
	mlbs := models.MLBGames{}
	err := mlbs.LoadByDateIncomplete(ctx, tp.Querier, date)
	if err != nil {
		return fmt.Errorf("error loading mlb for update: %w", err)
	}
	log.Printf("updating MLB games %d\n", len(mlbs))
	l.Info("updating mlb games", "amount", len(mlbs))

	for _, mlb := range mlbs {
		err := mlb.UpdateScore(ctx, tp.Querier, tp.client)
		if err != nil {
			l.Error("error updating mlb score", "error", err, "gameId", mlb.GameID)
		}
	}

	// Do same for NHL.
	nhls := models.NHLGames{}
	err = nhls.LoadByDateIncomplete(tp.Querier, date)
	if err != nil {
		return fmt.Errorf("error loading nhl for update: %w", err)
	}
	log.Printf("updating NHL games %d\n", len(nhls))
	for _, nhl := range nhls {
		if nhl.Status == "Final" {
			continue
		}
		err := nhl.UpdateScore(ctx, tp.client, tp.Querier)
		if err != nil {
			if errors.Is(err, models.ErrorNotFinished) {
				l.Debug("no game update", "msg", err)
			} else {
				return fmt.Errorf("error updating nhl score :%w", err)
			}
		}
	}
	return nil
}
