package main

import (
	"context"
	"fmt"
	"net/http"
)

// ContextKey defines context values.
type ContextKey string

// ContextUserKey is the username key.
const ContextUserKey ContextKey = "username"

// AuthHandler performs main authentication.
func AuthHandler(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, user, err := strategy.AuthenticateRequest(r)
		if err != nil {
			fmt.Println(err)
			code := http.StatusUnauthorized
			http.Error(w, http.StatusText(code), code)
			return
		}
		// Add username to context.
		ctx := context.WithValue(r.Context(), ContextUserKey, user.GetUserName())
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

// RoleHandler enforces user access to specific resources, or role-based access.
func RoleHandler(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		qUser, ok := r.URL.Query()["username"]
		if !ok {
			w.WriteHeader(406)
			w.Write([]byte("Must pass a username"))
			return
		}

		username := qUser[0]

		// Verify current-user has access to requested resource.
		authUser := r.Context().Value(ContextUserKey)
		if authUser != username {
			code := http.StatusUnauthorized
			http.Error(w, http.StatusText(code), code)
			return
		}

		next.ServeHTTP(w, r)
	})
}
