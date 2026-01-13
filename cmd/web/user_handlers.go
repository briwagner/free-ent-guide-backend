package main

import (
	"encoding/json"
	"fmt"
	"free-ent-guide-backend/models"
	"free-ent-guide-backend/models/modelstore"
	"net/http"
	"strconv"

	"github.com/shaj13/go-guardian/v2/auth"
)

// UsersCreateToken will return token for authenticated users.
// Responds to /v1/users/token
func (app *App) UsersCreateToken(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	username := r.Context().Value(ContextUserKey)

	// Load user to get ID and add to token.
	q := r.Context().Value(models.SqlcStorageContextKey).(*modelstore.Queries)
	u := &models.User{Email: fmt.Sprintf("%v", username)}
	err := u.FindByEmail(q)

	if err != nil {
		app.l.Error("error createToken", "error", err)
		w.WriteHeader(500)
		return
	}

	// @todo: add named scopes.

	jwtToken := app.Authy.IssueJWT(u.Email, fmt.Sprintf("%d", u.ID))
	// Does not throw error, only empty string, if it fails.
	if jwtToken == "" {
		app.l.Error("failed to generate JWT", "username", username)
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
func (app *App) UsersRevokeToken(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var formData UserTokenData
	err := decoder.Decode(&formData)
	if err != nil {
		app.l.Error("error decoding token", "error", err)
		panic(err)
	}

	if formData.Token == "" {
		w.Write([]byte("no data"))
		w.WriteHeader(400)
		return
	}

	err = auth.Revoke(app.JWTStrategy, formData.Token)
	if err != nil {
		app.l.Error("error revoking token", "error", err)
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
func (app *App) UsersCreate(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var formData UserCreateData
	err := decoder.Decode(&formData)
	if err != nil {
		app.l.Error("error usersCreate", "error", err)
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

	q := r.Context().Value(models.SqlcStorageContextKey).(*modelstore.Queries)
	user := &models.User{Email: formData.Username, Password: formData.Password}
	err = user.Create(q)
	if err != nil {
		app.l.Error("error usersCreate", "error", err)
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Error creating user. %v", err)))
		return
	}

	app.l.Info("user created", "username", user.Email)

	w.WriteHeader(204)
}

// UsersGetData returns the stored zip-codes for a user.
// Responds to /v1/users/get-zip?username={name}
func (app *App) UsersGetData(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	switch r.Method {
	case "OPTIONS":
		w.WriteHeader(200)
		return
	case "GET":
		// Infer user from authentication.
		uname := r.Context().Value(ContextUserKey)
		username := fmt.Sprintf("%v", uname)
		if username == "" {
			w.WriteHeader(404)
			w.Write([]byte("User not found"))
			return
		}

		q := r.Context().Value(models.SqlcStorageContextKey).(*modelstore.Queries)
		user := &models.User{Email: username}
		err := user.LoadData(q)
		if err != nil {
			w.WriteHeader(404)
			w.Write([]byte("Not found"))
			return
		}

		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			app.l.Error("error usersGetData", "error", err)
		}
		return
	default:
		w.WriteHeader(405)
		w.Write([]byte("Only GET requests are accepted."))
	}
}

// UsersAddZip creates or appends a zip code to user in storage.
// Responds to /v1/users/add-zip?username={name}&zip={zipcode}
func (app *App) UsersAddZip(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	switch r.Method {
	case "OPTIONS":
		w.WriteHeader(200)
		return
	case "POST":
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

		q := r.Context().Value(models.SqlcStorageContextKey).(*modelstore.Queries)
		user := &models.User{Email: username}
		newZip, err := strconv.ParseInt(qZip, 10, 64)
		if err != nil {
			app.l.Error("error parsing zip", "error", err)
			w.WriteHeader(500)
			return
		}
		err = user.AddZip(q, newZip)
		if err != nil {
			app.l.Error("storage error", "error", err)
			w.WriteHeader(500)
			return
		}

		err = user.GetZips(q)
		if err != nil {
			app.l.Error("error getZips", "error", err)
			w.WriteHeader(500)
		}

		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			app.l.Error("error usersGetZip", "error", err)
		}
		return
	default:
		w.WriteHeader(405)
		w.Write([]byte("Only POST requests are accepted."))
	}
}

// UsersDeleteZip removes a zip-code from a users list.
// Responds to /v1/users/delete-zip?username={name}&zip={zipcode}
func (app *App) UsersDeleteZip(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	switch r.Method {
	case "OPTIONS":
		w.WriteHeader(200)
		return
	case "POST":
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

		q := r.Context().Value(models.SqlcStorageContextKey).(*modelstore.Queries)
		user := &models.User{Email: username}
		delZip, err := strconv.ParseInt(qZip, 10, 64)
		if err != nil {
			app.l.Error("error parsing zip", "error", err)
			w.WriteHeader(500)
			return
		}
		err = user.DeleteZip(q, delZip)
		if err != nil {
			app.l.Error("error deleteZip", "error", err)
			w.WriteHeader(500)
			w.Write([]byte("Error deleting data."))
			return
		}

		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			app.l.Error("error deleteZip", "error", err)
		}
		return
	default:
		w.WriteHeader(405)
		w.Write([]byte("Only POST requests are accepted."))
	}
}

// UsersClearZip removes all stored zip-codes.
// Responds to /v1/users/clear-zip?username={name}
func (app *App) UsersClearZip(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	switch r.Method {
	case "POST":

		// Infer user from authentication.
		uname := r.Context().Value(ContextUserKey)
		username := fmt.Sprintf("%v", uname)
		if username == "" {
			w.WriteHeader(404)
			w.Write([]byte("User not found"))
			return
		}

		q := r.Context().Value(models.SqlcStorageContextKey).(*modelstore.Queries)
		user := &models.User{Email: username}
		err := user.ClearZips(q)
		if err != nil {
			app.l.Error("error clearZip", "error", err)
			w.Write([]byte("Error removing from storage."))
			return
		}
		w.Write([]byte("ok"))
	default:
		w.WriteHeader(405)
		w.Write([]byte("Only POST requests are accepted."))
	}
}

// TV SHOWS

// UsersAddShow creates or appends to user's stored shows.
// Responds to /v1/users/add-show?show={show_id}
func (app *App) UsersAddShow(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	switch r.Method {
	case "OPTIONS":
		w.WriteHeader(200)
		return
	case "POST":
		// Infer user from authentication.
		uname := r.Context().Value(ContextUserKey)
		username := fmt.Sprintf("%v", uname) /// make it a string
		if username == "" {
			w.WriteHeader(404)
			w.Write([]byte("User not found"))
			return
		}
		qShow := r.URL.Query().Get("show")
		if qShow == "" {
			w.WriteHeader(406)
			w.Write([]byte("Must pass a show ID"))
			return
		}

		q := r.Context().Value(models.SqlcStorageContextKey).(*modelstore.Queries)
		user := &models.User{Email: username}
		newShow, err := strconv.ParseInt(qShow, 10, 64)
		if err != nil {
			app.l.Error("error parsing show", "error", err)
			w.WriteHeader(500)
			return
		}
		err = user.AddShow(q, newShow)
		if err != nil {
			app.l.Error("storage error", "error", err)
			w.WriteHeader(500)
			return
		}

		err = user.LoadData(q)
		if err != nil {
			app.l.Error("error usersAddShow", "error", err)
			w.WriteHeader(500)
		}

		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			app.l.Error("error addShow", "error", err)
		}
		return
	default:
		w.WriteHeader(405)
		w.Write([]byte("Only POST requests are accepted."))
	}
}

// UsersDeleteShow removes shows from user's saved shows.
// Responds to /v1/users-delete-show?show={show_id}
func (app *App) UsersDeleteShow(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	switch r.Method {
	case "OPTIONS":
		w.WriteHeader(200)
		return
	case "POST":
		// Infer user from authentication.
		uname := r.Context().Value(ContextUserKey)
		username := fmt.Sprintf("%v", uname) // make it a string
		if username == "" {
			w.WriteHeader(404)
			w.Write([]byte("User not found"))
			return
		}
		qShow := r.URL.Query().Get("show")
		if qShow == "" {
			w.WriteHeader(406)
			w.Write([]byte("Must pass a show ID"))
			return
		}

		q := r.Context().Value(models.SqlcStorageContextKey).(*modelstore.Queries)
		user := &models.User{Email: username}
		show, err := strconv.ParseInt(qShow, 10, 64)
		if err != nil {
			app.l.Error("error parsing show", "error", err)
			w.WriteHeader(500)
			return
		}

		err = user.DeleteShow(q, show)
		if err != nil {
			app.l.Error("storage error", "error", err)
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(http.StatusNoContent)
		return
	default:
		w.WriteHeader(405)
		w.Write([]byte("Only POST requests are accepted."))
	}
}
