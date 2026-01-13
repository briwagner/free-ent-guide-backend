package main

import (
	"context"
	"errors"
	"fmt"
	"free-ent-guide-backend/models/modelstore"
	"free-ent-guide-backend/pkg/cred"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"go.opentelemetry.io/otel/trace"
)

type (
	TaskRunner func(ctx context.Context, l *slog.Logger, tp TaskPayload, args []string) error

	TaskPayload struct {
		Cred    *cred.Cred
		Querier *modelstore.Queries
		client  *http.Client
	}

	Task struct {
		Command     string
		Description string
		Subcommands []string
		Args        int // count plus one (itself). TODO this is kinda dumb
		Runner      TaskRunner
	}

	TaskCommander struct {
		Cred    *cred.Cred
		Querier *modelstore.Queries
		Tasks   map[string]Task
		Client  *http.Client

		l *slog.Logger
		t trace.Tracer
	}
)

// Print shows the help message.
func (tc *TaskCommander) Print() string {
	var sb strings.Builder

	sb.WriteString("\nAvailable commands:\n\n")

	for _, cmd := range tc.Tasks {
		sb.WriteString(cmd.Command + ": " + cmd.Description + "\n")
		sb.WriteString("  Subcommands: " + strings.Join(cmd.Subcommands, ", ") + "\n")
		sb.WriteString("  Required arguments: " + strconv.Itoa(cmd.Args-1) + "\n")

		sb.WriteString("\n")
	}

	return sb.String()
}

func (tc *TaskCommander) Run(ctx context.Context, args []string) error {
	cmd := args[1]

	t, ok := tc.Tasks[cmd]
	if !ok {
		return errors.New("command not found: " + cmd)
	}

	if len(args)-1 < t.Args {
		return fmt.Errorf("%s command requires %d arguments. Try: %s", cmd, t.Args, strings.Join(t.Subcommands, ", "))
	}

	tp := TaskPayload{Cred: tc.Cred, Querier: tc.Querier, client: tc.Client}
	return t.Runner(ctx, tc.l, tp, args)
}
