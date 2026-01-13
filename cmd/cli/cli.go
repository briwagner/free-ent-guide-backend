package main

import (
	"context"
	"fmt"
	"free-ent-guide-backend/models"
	"free-ent-guide-backend/models/modelstore"
	"free-ent-guide-backend/pkg/bri_otel"
	"free-ent-guide-backend/pkg/cred"
	"log"
	"os"
	"strings"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/log/global"
)

var (
	format = "2006-01-02"
)

const (
	instrumentationName    = "github.com/briwagner/free-entertainment-guide-backend"
	instrumentationVersion = "0.1.0"
	appName                = "cli-ent-api"
)

func main() {
	// Set-up application config.
	var c cred.Cred
	c.GetCreds("creds", ".")

	tc := buildTaskCommander(&c, modelstore.New(models.Setup(c)))

	// No sub-command was passed.
	if len(os.Args) <= 1 {
		fmt.Println(tc.Print())
		return
	}

	// Setup OTel for logging, tracing.
	ctx := context.Background()
	lp, tp := bri_otel.SetupOtel(ctx, instrumentationName, instrumentationVersion, appName)
	if lp == nil || tp == nil {
		panic("failed to setup open telemetry")
	}
	defer func() {
		// TODO combine these shutdown funcs
		if err := lp.Shutdown(ctx); err != nil {
			log.Printf("shutdown logger error %v", err)
		}
		if err := tp.Shutdown(ctx); err != nil {
			log.Printf("shutdown tracer error %v", err)
		}
	}()
	global.SetLoggerProvider(lp)
	tc.l = otelslog.NewLogger(appName)
	tc.t = otel.Tracer(appName)

	err := tc.Run(ctx, os.Args)
	if err != nil {
		log.Print(err)
		tc.l.Error("error runner", " error", err, "commands", strings.Join(os.Args, ", "))
		return
	}
	tc.l.Info("task complete", "commands", strings.Join(os.Args, ", "))
	// If this is hanging and won't close, then Otel is not connected.
}

func buildTaskCommander(c *cred.Cred, q *modelstore.Queries) TaskCommander {
	tc := TaskCommander{
		Cred:    c,
		Querier: q,
		Client:  bri_otel.NewOtelClient(5),
	}

	tasks := make(map[string]Task)

	tasks["nhl"] = Task{
		Command:     "nhl",
		Description: "Fetch new games from NHL, for date provided",
		Args:        2,
		Subcommands: []string{"valid date", "'last'", "'latest'"},
		Runner:      handleNHL,
	}
	tasks["mlb"] = Task{
		Command:     "mlb",
		Description: "Fetch new games from MLB, for date provided",
		Args:        2,
		Subcommands: []string{"valid date", "'last'", "'latest'"},
		Runner:      handleMLB,
	}
	tasks["cache"] = Task{
		Command:     "cache",
		Description: "View or clear content cache",
		Args:        2,
		Subcommands: []string{"show", "stale", "clear", "wipe"},
		Runner:      handleCache,
	}
	tasks["gup"] = Task{
		Command:     "gup",
		Description: "Fetch game updates for NHL or MLB, for date provided",
		Subcommands: []string{"valid date, e.g. 2025-01-06"},
		Args:        2,
		Runner:      handleGameUpdates,
	}
	tasks["discover"] = Task{
		Command:     "discover",
		Description: "Fetch new movies, for data provided",
		Subcommands: []string{"valid date, e.g. 2025-01-06"},
		Args:        2,
		Runner:      handleDiscoverMovies,
	}
	tasks["backup"] = Task{
		Command:     "backup",
		Description: "Generate database backup file",
		Args:        1,
		Runner:      handleBackup,
	}
	tasks["user"] = Task{
		Command:     "user",
		Description: "Reset user data",
		Subcommands: []string{"reset data, e.g. reset joe@email.com"},
		Args:        3,
		Runner:      handleUser,
	}

	tc.Tasks = tasks
	return tc
}
