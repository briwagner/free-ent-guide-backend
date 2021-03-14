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
		if r.Method == "OPTIONS" {
			enableCors(&w)
			// TODO: above is not enough for preflight options. Add below to default action?
			(w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
			(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			w.WriteHeader(200)
			return
		}
		enableCors(&w)
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
		qUser := r.URL.Query().Get("username")
		if qUser == "" {
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
