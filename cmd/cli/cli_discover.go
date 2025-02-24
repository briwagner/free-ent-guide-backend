package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"free-ent-guide-backend/pkg/moviedb"
	"log"
	"os"
	"time"
)

func handleDiscoverMovies(tp TaskPayload, args []string) error {
	subCo := args[2]
	if subCo == "" {
		fmt.Println("Must pass date for discover movies lookup.")
		return nil
	}

	fmt.Println("Fetching discover movies...")

	_, err := time.Parse("2006-01-02", subCo)
	if err != nil {
		return err
	}

	mdb := moviedb.NewMovieDB(tp.Cred.Moviedb)
	results, err := mdb.GetDiscoverPaged(subCo)
	if err != nil {
		return err
	}

	if len(results) == 0 {
		return errors.New("no results for discovery movies")
	}

	fn := "discover.json"
	f, err := os.Create(fn)
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

	fmt.Printf("File generated: %s. Results %d \n", fn, len(results))

	b := "free-entertainment-guide.com"
	log.Printf("pushing to bucket %s", b)
	return tp.Cred.Spaces.PutFile("discover.json", "discover.json", b)
}
