package main

import (
	"encoding/json"
	"fmt"
	"free-ent-guide-backend/models"
	"log"
	"net/http"

	"github.com/shaj13/go-guardian/v2/auth"
)

// UsersCreateToken will return token for authenticated users.
// Responds to /v1/users/token
func UsersCreateToken(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	username := r.Context().Value(ContextUserKey)

	// Load user to get ID and add to token.
	DB := r.Context().Value(models.StorageContextKey).(models.Store)
	u := models.User{Name: fmt.Sprintf("%v", username)}
	err := u.LoadByName(DB)
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
		return
	}

	// @todo: add named scopes.

	jwtToken := author.IssueJWT(u.Name, fmt.Sprintf("%d", u.ID))
	// Does not throw error, only empty string, if it fails.
	if jwtToken == "" {
		log.Printf("Failed to generate JWT for %s", username)
		w.WriteHeader(500)
		return
	}

	body := fmt.Sprintf("%s\n", jwtToken)
	w.Write([]byte(body))
}

// UserTokenData is a map to accept form data for user-token-revoke.
type UserTokenData struct {
	Token string
}

// UsersRevokeToken revokes the token.
// Responds to /v1/users/token
func UsersRevokeToken(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		enableCors(&w)
		// TODO: above is not enough for preflight options. Add below to default action?
		(w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.WriteHeader(200)
		return
	}

	enableCors(&w)

	decoder := json.NewDecoder(r.Body)
	var formData UserTokenData
	err := decoder.Decode(&formData)
	if err != nil {
		panic(err)
	}

	if formData.Token == "" {
		w.Write([]byte("no data"))
		w.WriteHeader(400)
		return
	}

	err = auth.Revoke(jwtStrategy, formData.Token)
	if err != nil {
		w.Write([]byte("error revoking token"))
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(204)
}

// UserCreateData is a map to accept form data for user-create form.
type UserCreateData struct {
	Username string
	Password string
}

// UsersCreate adds a user to storage.
// Responds to /v1/users/create
func UsersCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		enableCors(&w)
		// TODO: above is not enough for preflight options. Add below to default action?
		(w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.WriteHeader(200)
		return
	}

	enableCors(&w)

	decoder := json.NewDecoder(r.Body)
	var formData UserCreateData
	err := decoder.Decode(&formData)
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
		return
	}

	// Get data from form-data.
	if formData.Username == "" {
		w.WriteHeader(406)
		w.Write([]byte("Must pass a username"))
		return
	}

	if formData.Password == "" {
		w.WriteHeader(406)
		w.Write([]byte("Must pass a password"))
		return
	}

	DB := r.Context().Value(models.StorageContextKey).(models.Store)
	user := &models.User{Name: formData.Username, Password: formData.Password}
	err = user.Create(DB)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Error creating user. %v", err)))
		return
	}

	w.WriteHeader(204)
}

// UsersGetZip returns the stored zip-codes for a user.
// Responds to /v1/users/get-zip?username={name}
func UsersGetZip(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	switch r.Method {
	case "OPTIONS":
		enableCors(&w)
		// TODO: above is not enough for preflight options. Add below to default action?
		(w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.WriteHeader(200)
		return
	case "GET":
		enableCors(&w)

		// Infer user from authentication.
		uname := r.Context().Value(ContextUserKey)
		username := fmt.Sprintf("%v", uname)
		if username == "" {
			w.WriteHeader(404)
			w.Write([]byte("User not found"))
			return
		}

		DB := r.Context().Value(models.StorageContextKey).(models.Store)
		user := &models.User{Name: username}
		err := user.GetZips(DB)
		if err != nil {
			w.WriteHeader(404)
			w.Write([]byte("Not found"))
			return
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
	case "OPTIONS":
		enableCors(&w)
		// TODO: above is not enough for preflight options. Add below to default action?
		(w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.WriteHeader(200)
		return
	case "POST":
		enableCors(&w)

		// Infer user from authentication.
		uname := r.Context().Value(ContextUserKey)
		username := fmt.Sprintf("%v", uname)
		if username == "" {
			w.WriteHeader(404)
			w.Write([]byte("User not found"))
			return
		}
		qZip := r.URL.Query().Get("zip")
		if qZip == "" {
			w.WriteHeader(406)
			w.Write([]byte("Must pass a zip"))
			return
		}

		DB := r.Context().Value(models.StorageContextKey).(models.Store)
		user := &models.User{Name: username}
		err := user.AddZip(DB, qZip)
		if err != nil {
			log.Printf("Storage error %s", err.Error())
			w.WriteHeader(500)
			w.Write([]byte("Error storing zip."))
			return
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
		w.Write([]byte("Only POST requests are accepted."))
	}
}

// UsersDeleteZip removes a zip-code from a users list.
// Responds to /v1/users/delete-zip?username={name}&zip={zipcode}
func UsersDeleteZip(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "OPTIONS":
		enableCors(&w)
		// TODO: above is not enough for preflight options. Add below to default action?
		(w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.WriteHeader(200)
		return
	case "POST":
		enableCors(&w)

		// Infer user from authentication.
		uname := r.Context().Value(ContextUserKey)
		username := fmt.Sprintf("%v", uname)
		if username == "" {
			w.WriteHeader(404)
			w.Write([]byte("User not found"))
			return
		}
		qZip := r.URL.Query().Get("zip")
		if qZip == "" {
			w.WriteHeader(406)
			w.Write([]byte("Must pass a zip"))
			return
		}

		DB := r.Context().Value(models.StorageContextKey).(models.Store)
		user := &models.User{Name: username}
		err := user.DeleteZip(DB, qZip)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Error deleting data."))
			return
		}

		userJSON, err := json.Marshal(user)
		if err != nil {
			log.Print("Error marshaling user zips; UsersDeleteZip")
			w.WriteHeader(200)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(userJSON))
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

		// Infer user from authentication.
		uname := r.Context().Value(ContextUserKey)
		username := fmt.Sprintf("%v", uname)
		if username == "" {
			w.WriteHeader(404)
			w.Write([]byte("User not found"))
			return
		}

		DB := r.Context().Value(models.StorageContextKey).(models.Store)
		user := &models.User{Name: username}
		err := user.ClearZips(DB)
		if err != nil {
			w.Write([]byte("Error removing from storage."))
			return
		}
		w.Write([]byte("ok"))
	default:
		w.WriteHeader(405)
		w.Write([]byte("Only POST requests are accepted."))
	}
}
