package main

import (
	"encoding/json"
	"fmt"
	"free-ent-guide-backend/models"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/shaj13/go-guardian/v2/auth"
)

// UsersCreateToken will return token for authenticated users.
// Responds to /v1/users/token
func UsersCreateToken(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value(ContextUserKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "ent_app",
		"sub": username,
		"aud": "any",
		"exp": time.Now().Add(time.Minute * time.Duration(c.TokenDuration)).Unix(),
	})

	jwtToken, _ := token.SignedString([]byte(c.TokenSecret))

	user := auth.NewDefaultUser(fmt.Sprintf("%s", username), "1", nil, nil)
	auth.Append(tokenStrategy, jwtToken, user)
	body := fmt.Sprintf("token: %s \n", jwtToken)
	w.Write([]byte(body))
}

// UsersCreate adds a user to storage.
// Responds to /v1/users/create
func UsersCreate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		qUser := r.URL.Query().Get("username")
		if qUser == "" {
			w.WriteHeader(406)
			w.Write([]byte("Must pass a username"))
			return
		}

		qPass := r.URL.Query().Get("password")
		if qPass == "" {
			w.WriteHeader(406)
			w.Write([]byte("Must pass a password"))
			return
		}

		user := &models.User{Name: qUser, Password: qPass}
		err := user.Create(cacheClient)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(fmt.Sprintf("Error creating user. %v", err)))
			return
		}

		w.Write([]byte("ok"))
	default:
		w.WriteHeader(405)
		w.Write([]byte("Only POST requests are accepted."))
	}
}

// UsersCreateToken returns a token if user is authenticated.
// Responds to /v1/users/token

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
		qUser := r.URL.Query().Get("username")
		if qUser == "" {
			w.WriteHeader(406)
			w.Write([]byte("Must pass a username"))
			return
		}

		records, err := cacheClient.LRange(fmt.Sprintf("user:%s:zip", qUser), 0, -1).Result()
		if err != nil || len(records) == 0 {
			w.WriteHeader(404)
			w.Write([]byte("Not found"))
			return
		}

		user := &models.User{
			Name: qUser,
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
		qUser := r.URL.Query().Get("username")
		if qUser == "" {
			w.WriteHeader(406)
			w.Write([]byte("Must pass a username"))
			return
		}
		qZip := r.URL.Query().Get("zip")
		if qZip == "" {
			w.WriteHeader(406)
			w.Write([]byte("Must pass a zip"))
			return
		}

		res := cacheClient.RPush(fmt.Sprintf("user:%s:zip", qUser), qZip)
		if res.Err() != nil {
			log.Printf("Storage error %v", res.Err())
			w.WriteHeader(500)
			w.Write([]byte("Error storing zip."))
			return
		}

		// Reload data to return to client.
		records, err := cacheClient.LRange(fmt.Sprintf("user:%s:zip", qUser), 0, -1).Result()
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
		qUser := r.URL.Query().Get("username")
		if qUser == "" {
			w.WriteHeader(406)
			w.Write([]byte("Must pass a username"))
			return
		}
		qZip := r.URL.Query().Get("zip")
		if qZip == "" {
			w.WriteHeader(406)
			w.Write([]byte("Must pass a zip"))
			return
		}

		res, err := cacheClient.LRem(fmt.Sprintf("user:%s:zip", qUser), 0, qZip).Result()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Error deleting data."))
			return
		}

		// TODO: restrict entering repeated values. Then change this to == 1.
		// Redis returns # of items removed.
		if res >= 1 {
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

		qUser := r.URL.Query().Get("username")
		if qUser == "" {
			w.Write([]byte("Must pass a username"))
			return
		}
		res := cacheClient.Del(fmt.Sprintf("user:%s:zip", qUser))
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
