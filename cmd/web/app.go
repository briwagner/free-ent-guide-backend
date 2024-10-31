package main

import (
	"free-ent-guide-backend/pkg/authenticator"
	"time"

	"github.com/shaj13/go-guardian/v2/auth"
	"github.com/shaj13/go-guardian/v2/auth/strategies/basic"
	"github.com/shaj13/go-guardian/v2/auth/strategies/jwt"
	"github.com/shaj13/go-guardian/v2/auth/strategies/union"

	"github.com/shaj13/libcache"
)

type App struct {
	Authy       *authenticator.Authenticator
	JWTStrategy auth.Strategy
	Strategy    union.Union
}

// Define strategies, set token expiration time.
func (a *App) setupGoGuardian() {
	author := authenticator.Authenticator{
		Audience: jwt.SetAudience("any"),
		Issuer:   jwt.SetIssuer("ent_api"),
		Duration: c.TokenDuration,
	}
	author.AttachSecret(c.TokenSecret)

	cacheObj := libcache.FIFO.New(0)
	cacheObj.SetTTL(time.Hour * 72)
	// TODO this is deprecated
	// cacheObj.RegisterOnExpired(func(key, _ interface{}) {
	// 	cacheObj.Peek(key) // TODO this is pointless.
	// })

	jwtStrategy := jwt.New(cacheObj, author.Secret, author.Audience)
	basicStrategy := basic.NewCached(validateUser, cacheObj)
	author.Strategy = union.New(jwtStrategy, basicStrategy)

	// TODO this is for revoke. Can we not set this?
	a.JWTStrategy = jwtStrategy

	a.Authy = &author
}
