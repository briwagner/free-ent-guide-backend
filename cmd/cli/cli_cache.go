package main

import (
	"fmt"
	"free-ent-guide-backend/models"
	"free-ent-guide-backend/models/modelstore"
	"log"
	"slices"
	"strings"
	"time"
)

func handleCache(q *modelstore.Queries, args []string) {
	op := args[2]
	ops := []string{"show", "stale", "clear", "wipe"}
	if !slices.Contains(ops, op) {
		fmt.Printf("For 'cache', must pass one of %s\n", strings.Join(ops, ", "))
	}
	cs := &models.Caches{}

	if op == "show" {
		err := cs.ShowAll(q)
		if err != nil {
			log.Print(err)
			return
		}
		if len(*cs) == 0 {
			log.Print("no records found")
			return
		}
		for _, c := range *cs {
			fmt.Println(c)
		}
		return
	}

	if op == "stale" {
		err := cs.ShowStale(time.Now(), q)
		if err != nil {
			log.Print(err)
			return
		}
		if len(*cs) == 0 {
			log.Print("no records found")
			return
		}
		for _, c := range *cs {
			fmt.Println(c)
		}
		return
	}

	if op == "clear" {
		rows, err := cs.DropStale(time.Now(), q)
		if err != nil {
			log.Print(err)
			return
		}
		log.Printf("Dropped %d records.\n", rows)
	}

	if op == "wipe" {
		rows, err := cs.DropAll(q)
		if err != nil {
			log.Print(err)
			return
		}
		log.Printf("%d caches wiped\n", rows)
	}
}
