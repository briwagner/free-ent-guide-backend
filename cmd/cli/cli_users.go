package main

import (
	"context"
	"fmt"
	"free-ent-guide-backend/models"
	"log/slog"
	"slices"
	"strings"
)

func handleUser(ctx context.Context, l *slog.Logger, tp TaskPayload, args []string) error {
	op := args[2]
	ops := []string{"reset"}
	if !slices.Contains(ops, op) {
		return fmt.Errorf("for 'user', must pass one of %s", strings.Join(ops, ", "))
	}

	email := args[3]
	if email == "" {
		return fmt.Errorf("must pass email to reset user data")
	}

	// If we query for user when the column is bad, it will fail.
	// So check for user exists after this step.
	u := models.User{Email: email}
	err := u.FixData(ctx, tp.Querier)
	if err != nil {
		return err
	}
	l.Info("user reset", "username", u.Email)
	// Now return if user is valid.
	return u.FindByEmail(tp.Querier)
}
