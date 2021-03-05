package main

import (
	"encoding/json"
	"fmt"
	"free-ent-guide-backend/models"
	"log"
	"net/http"
)

// UsersGetZip returns the stored zip-codes for a user.
// Responds to /v1/users/get-zip?username={name}
func UsersGetZip(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		enableCors(&w)
		if c.Cache != true {
			w.WriteHeader(500)
			w.Write([]byte("Database not found."))
			return
		}
		qUser, ok := r.URL.Query()["username"]
		if !ok {
			w.WriteHeader(406)
			w.Write([]byte("Must pass a username"))
			return
		}

		username := qUser[0]
		records, err := cacheClient.LRange(fmt.Sprintf("user:%s", username), 0, -1).Result()
		if err != nil {
			w.WriteHeader(404)
			w.Write([]byte("User not found"))
			return
		}

		user := &models.User{
			Name: username,
			Zips: records,
		}
		jsonU, err := json.Marshal(user)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("JSON error"))
			return
		}
		w.Write([]byte(jsonU))
	default:
		w.WriteHeader(405)
		w.Write([]byte("Only GET requests are accepted."))
	}
}

// UsersAddZip creates or appends a zip code to user in storage.
// Responds to /v1/users/add-zip?username={name}&zip={zipcode}
func UsersAddZip(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		enableCors(&w)
		if c.Cache != true {
			w.Write([]byte("Database not found."))
			return
		}

		// Get parameters from request.
		qUser, ok := r.URL.Query()["username"]
		if !ok {
			w.WriteHeader(406)
			w.Write([]byte("Must pass a username"))
			return
		}
		qZip, ok := r.URL.Query()["zip"]
		if !ok {
			w.WriteHeader(406)
			w.Write([]byte("Must pass a zip"))
			return
		}

		u := qUser[0]
		z := qZip[0]

		// User:{name} is a List type.
		res := cacheClient.RPush(fmt.Sprintf("user:%s", u), z)
		if res.Err() != nil {
			log.Printf("Storage error %v", res.Err())
			w.WriteHeader(500)
			w.Write([]byte("Error storing zip."))
			return
		}

		// Reload data to return to client.
		records, err := cacheClient.LRange(fmt.Sprintf("user:%s", u), 0, -1).Result()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Error fetching new data."))
			return
		}

		user := &models.User{Zips: records}
		jsonU, err := json.Marshal(user)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("JSON error"))
			return
		}
		w.Write([]byte(jsonU))
	default:
		w.WriteHeader(405)
		w.Write([]byte("Only POST requests are accepted."))
	}
}

// UsersClearZip removes all stored zip-codes.
// Responds to /v1/users/clear-zip?username={name}
func UsersClearZip(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		enableCors(&w)
		if c.Cache != true {
			w.Write([]byte("Database not found."))
			return
		}

		qUser, ok := r.URL.Query()["username"]
		if !ok {
			w.Write([]byte("Must pass a username"))
			return
		}
		res := cacheClient.Del(fmt.Sprintf("user:%s", qUser[0]))
		if res.Err() != nil {
			w.Write([]byte("Error removing from storage."))
			return
		}
		w.Write([]byte("OK"))
	default:
		w.WriteHeader(405)
		w.Write([]byte("Only POST requests are accepted."))
	}
}
