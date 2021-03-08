package main

import (
	"encoding/json"
	"fmt"
	"free-ent-guide-backend/models"
	"log"
	"net/http"
)

// UsersCreate adds a user to storage.
// Responds to /v1/users/create
func UsersCreate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		qUser, ok := r.URL.Query()["username"]
		if !ok {
			w.WriteHeader(406)
			w.Write([]byte("Must pass a username"))
			return
		}

		qPass, ok := r.URL.Query()["password"]
		if !ok {
			w.WriteHeader(406)
			w.Write([]byte("Must pass a password"))
			return
		}

		user := &models.User{Name: qUser[0], Password: qPass[0]}
		err := user.Create(cacheClient)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(fmt.Sprintf("Error creating user. %v", err)))
			return
		}

		w.Write([]byte("ok"))
	default:
		w.WriteHeader(405)
		w.Write([]byte("Only GET requests are accepted."))
	}
}

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
		records, err := cacheClient.LRange(fmt.Sprintf("user:%s:zip", username), 0, -1).Result()
		if err != nil || len(records) == 0 {
			w.WriteHeader(404)
			w.Write([]byte("Not found"))
			return
		}

		user := &models.User{
			Name: username,
			Zips: records,
		}
		userJSON, err := json.Marshal(user)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("JSON error"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(userJSON))
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

		res := cacheClient.RPush(fmt.Sprintf("user:%s:zip", u), z)
		if res.Err() != nil {
			log.Printf("Storage error %v", res.Err())
			w.WriteHeader(500)
			w.Write([]byte("Error storing zip."))
			return
		}

		// Reload data to return to client.
		records, err := cacheClient.LRange(fmt.Sprintf("user:%s:zip", u), 0, -1).Result()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Error fetching new data."))
			return
		}

		user := &models.User{Zips: records}
		userJSON, err := json.Marshal(user)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("JSON error"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(userJSON))
	default:
		w.WriteHeader(405)
		w.Write([]byte("Only POST requests are accepted."))
	}
}

// UsersDeleteZip removes a zip-code from a users list.
// Responds to /v1/users/delete-zip?username={name}&zip={zipcode}
func UsersDeleteZip(w http.ResponseWriter, r *http.Request) {
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

		res, err := cacheClient.LRem(fmt.Sprintf("user:%s:zip", u), 0, z).Result()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Error deleting data."))
			return
		}

		if res == 1 {
			w.Write([]byte("ok"))
		} else {
			w.Write([]byte("not found"))
		}
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
		res := cacheClient.Del(fmt.Sprintf("user:%s:zip", qUser[0]))
		if res.Err() != nil {
			w.Write([]byte("Error removing from storage."))
			return
		}
		w.Write([]byte("ok"))
	default:
		w.WriteHeader(405)
		w.Write([]byte("Only POST requests are accepted."))
	}
}
