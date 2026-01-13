package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"free-ent-guide-backend/pkg/moviedb"
	"log"
	"log/slog"
	"os"
	"time"
)

const (
	FN     = "discover.json"
	Domain = "free-entertainment-guide.com"
)

func handleDiscoverMovies(ctx context.Context, l *slog.Logger, tp TaskPayload, args []string) error {
	subCo := args[2]
	if subCo == "" {
		fmt.Println("Must pass date for discover movies lookup.")
		return nil
	}

	_, err := time.Parse("2006-01-02", subCo)
	if err != nil {
		return fmt.Errorf("error parsing date for %s: %w", subCo, err)
	}

	log.Printf("Fetching discover movies for %s...\n", subCo)

	mdb := moviedb.NewMovieDB(tp.Cred.Moviedb)
	results, err := mdb.GetDiscoverPaged(ctx, subCo)
	if err != nil {
		return fmt.Errorf("error doing GetDiscoverPaged: %w", err)
	}

	if len(results) == 0 {
		return errors.New("no results for discovery movies")
	}

	f, err := os.Create(FN)
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := json.Marshal(results)
	if err != nil {
		return err
	}

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	log.Printf("File generated: %s. Results %d \n", FN, len(results))
	l.Info("discover file generated", "filename", FN, "results", len(results))

	log.Printf("pushing to bucket %s", Domain)
	return tp.Cred.Spaces.PutFile(FN, FN, Domain)
}
