package main

import (
	"fmt"
	"free-ent-guide-backend/models"
	"log"
	"slices"
	"strings"
	"time"
)

func handleCache(tp TaskPayload, args []string) error {
	op := args[2]
	ops := []string{"show", "stale", "clear", "wipe"}
	if !slices.Contains(ops, op) {
		fmt.Printf("For 'cache', must pass one of %s\n", strings.Join(ops, ", "))
	}
	cs := &models.Caches{}

	if op == "show" {
		err := cs.ShowAll(tp.Querier)
		if err != nil {
			return err
		}
		if len(*cs) == 0 {
			log.Print("no records found")
			return nil
		}
		for _, c := range *cs {
			fmt.Println(c)
		}
		return nil
	}

	if op == "stale" {
		err := cs.ShowStale(time.Now(), tp.Querier)
		if err != nil {
			return err
		}
		if len(*cs) == 0 {
			log.Print("no records found")
			return nil
		}
		for _, c := range *cs {
			fmt.Println(c)
		}
		return nil
	}

	if op == "clear" {
		rows, err := cs.DropStale(time.Now(), tp.Querier)
		if err != nil {
			return err
		}
		log.Printf("Dropped %d records.\n", rows)
	}

	if op == "wipe" {
		rows, err := cs.DropAll(tp.Querier)
		if err != nil {
			return err
		}
		log.Printf("%d caches wiped\n", rows)
	}

	return nil
}
