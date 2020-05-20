package jwtutils

import (
	"net/http"
	"time"

	"goutils/strutils"

	"github.com/dgrijalva/jwt-go"
)

var (
	defaultAuth   TokenAuth
	ServerSecret  = strutils.RandomANI(32)
	TokenExpireIn = 2 * time.Hour
)

func init() {
	defaultAuth = TokenAuth{ServerSecret, &TokenExpireIn}
}

func NewTokenAuth() TokenAuth {
	tokenExpireIn := TokenExpireIn
	return TokenAuth{ServerSecret, &tokenExpireIn}
}

func IssueMapToken(w http.ResponseWriter, m map[string]interface{}) {
	defaultAuth.IssueMapToken(w, m)
}

func IssueStandardToken(w http.ResponseWriter, aud, sub string) {
	defaultAuth.IssueStandardToken(w, aud, sub)
}

func IsTokenValid(r *http.Request) bool {
	return defaultAuth.IsTokenValid(r)
}

func GetTokenClaims(r *http.Request) (jwt.MapClaims, error) {
	return defaultAuth.GetTokenClaims(r)
}
