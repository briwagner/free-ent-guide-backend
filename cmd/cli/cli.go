package main

import (
	"errors"
	"fmt"
	"free-ent-guide-backend/models"
	"free-ent-guide-backend/models/modelstore"
	"free-ent-guide-backend/pkg/cred"
	"log"
	"os"
	"strings"
)

var (
	format = "2006-01-02"
)

func main() {
	// Set-up application config.
	var c cred.Cred
	c.GetCreds("creds", ".")

	tc := buildTaskCommander(&c, modelstore.New(models.Setup(c)))

	if len(os.Args) <= 1 {
		fmt.Println(tc.Print())
		return
	}

	err := tc.Run(os.Args)
	if err != nil {
		log.Print(err)
	}
}

type (
	TaskRunner func(tp TaskPayload, args []string) error

	TaskPayload struct {
		Cred    *cred.Cred
		Querier *modelstore.Queries
	}

	Task struct {
		Command     string
		Description string
		Subcommands []string
		Args        int // count
		Runner      TaskRunner
	}

	TaskCommander struct {
		Cred    *cred.Cred
		Querier *modelstore.Queries
		Tasks   map[string]Task
	}
)

func (tc *TaskCommander) Print() string {
	cmds := make([]string, len(tc.Tasks))
	count := 0
	for cmdName := range tc.Tasks {
		cmds[count] = cmdName
		count++
	}

	return fmt.Sprintf(`
	Available commands: %s.
	'nhl', 'mlb' or 'gup' along with a date e.g. 2022-11-21.
	Or 'nhl', 'mlb' with 'last'.
	Or 'cache' with one of: show, stale, clear, wipe.
	Or 'discover' with date.
	Or 'backup'.

	`, strings.Join(cmds, ", "))
}

func (tc *TaskCommander) Run(args []string) error {
	cmd := args[1]

	t, ok := tc.Tasks[cmd]
	if !ok {
		return errors.New("command not found: " + cmd)
	}

	if len(args)-1 < t.Args {
		return fmt.Errorf("%s command requires %d arguments. Try: %s", cmd, t.Args, strings.Join(t.Subcommands, ", "))
	}

	tp := TaskPayload{Cred: tc.Cred, Querier: tc.Querier}
	return t.Runner(tp, args)
}

func buildTaskCommander(c *cred.Cred, q *modelstore.Queries) TaskCommander {
	tc := TaskCommander{
		Cred:    c,
		Querier: q,
	}

	tasks := make(map[string]Task)

	tasks["nhl"] = Task{
		Command:     "nhl",
		Description: "Fetch new games from NHL, for date provided",
		Args:        2,
		Subcommands: []string{"valid date", "'last'"},
		Runner:      handleNHL,
	}
	tasks["mlb"] = Task{
		Command:     "mlb",
		Description: "Fetch new games from MLB, for date provided",
		Args:        2,
		Subcommands: []string{"valid date", "'last'"},
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
		Subcommands: []string{"valid date, e.g. 202-01-06"},
		Args:        2,
		Runner:      handleGameUpdates,
	}
	tasks["discover"] = Task{
		Command:     "discover",
		Description: "Fetch new movies, for data provided",
		Subcommands: []string{"valid date, e.g. 202-01-06"},
		Args:        2,
		Runner:      handleDiscoverMovies,
	}
	tasks["backup"] = Task{
		Command:     "backup",
		Description: "Generate database backup file",
		Args:        1,
		Runner:      handleBackup,
	}

	tc.Tasks = tasks
	return tc
}
