package main

import (
	"encoding/json"
	"fmt"
	"free-ent-guide-backend/pkg/cred"
	"free-ent-guide-backend/pkg/moviedb"
	"log"
	"os"
	"time"
)

func handleDiscoverMovies(c cred.Cred, args []string) error {
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

	mdb := moviedb.MovieDb{Key: c.Moviedb}
	results, err := mdb.GetDiscoverPaged(subCo)
	if err != nil {
		return err
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
	return c.Spaces.PutFile("discover.json", "discover.json", b)
}
