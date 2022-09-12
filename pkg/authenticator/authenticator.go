package authenticator

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/shaj13/go-guardian/v2/auth"
	"github.com/shaj13/go-guardian/v2/auth/strategies/jwt"
	"github.com/shaj13/libcache"

	_ "github.com/shaj13/libcache/lru"
)

type Authenticator struct {
	Strategy  auth.Strategy
	Audience  auth.Option
	Issuer    auth.Option
	Duration  int64
	Secret    jwt.StaticSecret
	SecretKey string
	Cache     libcache.Cache
}

// SetSecretKey is a setter
func (a *Authenticator) SetSecretKey(s string) {
	a.SecretKey = s
}

// GetSecretKey is a getter
func (a *Authenticator) GetSecretKey() string {
	return a.SecretKey
}

// AttachSecret stores a secret
func (a *Authenticator) AttachSecret(s string) {
	secret := jwt.StaticSecret{
		ID:        "jwt_secret",
		Algorithm: jwt.HS256,
		Secret:    []byte(s),
	}
	a.Secret = secret
}

// IssueJWT generates a JWT with specified audience.
func (a *Authenticator) IssueJWT(username string, uid string) string {
	// Some user info is passed
	u := auth.NewUserInfo(username, uid, nil, nil)

	exp := jwt.SetExpDuration(time.Minute * time.Duration(a.Duration))
	// @todo: add jwt.SetNamedScopes()
	token, err := jwt.IssueAccessToken(u, a.Secret, a.Audience, a.Issuer, exp)
	if err != nil {
		log.Println(err)
		return ""
	}

	return token
}

// IssueToken generates a token.
// @todo: we're not using this.
func (a *Authenticator) IssueToken(username string, uid string) string {
	// Some user info is passed
	u := auth.NewUserInfo(username, uid, nil, nil)

	a.Cache.Store("token", u)

	token, err := jwt.IssueAccessToken(u, a.Secret)
	if err != nil {
		log.Println(err)
		return ""
	}

	return token
}

// HandleUserRequest manages a user request using auth
func (a *Authenticator) HandleUserRequest(token string, r *http.Request) (auth.Info, error) {
	user, err := a.Strategy.Authenticate(r.Context(), r)
	if err != nil {
		fmt.Printf("got error %s", err.Error())
		return nil, err
	}
	return user, nil
}
