package main

import (
	"context"
	"fmt"
	"free-ent-guide-backend/models"
	"log"
	"log/slog"
	"slices"
	"strings"
	"time"
)

func handleCache(ctx context.Context, l *slog.Logger, tp TaskPayload, args []string) error {
	op := args[2]
	ops := []string{"show", "stale", "clear", "wipe"}
	if !slices.Contains(ops, op) {
		fmt.Printf("For 'cache', must pass one of %s\n", strings.Join(ops, ", "))
	}
	cs := &models.Caches{}

	if op == "show" {
		err := cs.ShowAll(ctx, tp.Querier)
		if err != nil {
			return err
		}
		if len(*cs) == 0 {
			l.Info("no cache records found", "sub-command", op)
			return nil
		}
		for _, c := range *cs {
			fmt.Println(c)
		}
		return nil
	}

	if op == "stale" {
		err := cs.ShowStale(ctx, time.Now(), tp.Querier)
		if err != nil {
			return err
		}
		if len(*cs) == 0 {
			l.Info("no cache records found", "sub-command", op)
			return nil
		}
		for _, c := range *cs {
			fmt.Println(c)
		}
		return nil
	}

	if op == "clear" {
		rows, err := cs.DropStale(ctx, time.Now(), tp.Querier)
		if err != nil {
			return err
		}
		l.Info("dropped cache records", "rows", rows)
		fmt.Printf("dropped %d records\n", rows)
	}

	if op == "wipe" {
		rows, err := cs.DropAll(ctx, tp.Querier)
		if err != nil {
			return err
		}
		log.Printf("%d caches wiped\n", rows)
		l.Info("wiped cache records", "rows", rows)
	}

	return nil
}
