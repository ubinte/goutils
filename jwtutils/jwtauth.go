package jwtutils

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenAuth struct {
	ServerSecret  string
	TokenExpireIn *time.Duration
}

func (auth *TokenAuth) IssueMapToken(w http.ResponseWriter, m map[string]interface{}) {
	claims := make(jwt.MapClaims)
	for k, v := range m {
		claims[k] = v
	}
	auth.issueToken(w, claims)
}

func (auth *TokenAuth) IssueStandardToken(w http.ResponseWriter, aud, sub string) {
	claims := jwt.StandardClaims{
		Audience:  aud,
		Subject:   sub,
		ExpiresAt: time.Now().Add(*auth.TokenExpireIn).Unix(),
	}
	auth.issueToken(w, claims)
}

func (auth *TokenAuth) issueToken(w http.ResponseWriter, claims jwt.Claims) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	access_token, _ := token.SignedString([]byte(auth.ServerSecret))

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    access_token,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   int(auth.TokenExpireIn.Seconds()),
	})
}

func (auth *TokenAuth) IsTokenValid(r *http.Request) bool {
	if _, err := auth.getToken(r); err != nil {
		return false
	} else {
		return true
	}
}

func (auth *TokenAuth) GetTokenClaims(r *http.Request) (jwt.MapClaims, error) {
	if token, err := auth.getToken(r); err != nil {
		return nil, err
	} else {
		return token.Claims.(jwt.MapClaims), nil
	}
}

func (auth *TokenAuth) getToken(r *http.Request) (*jwt.Token, error) {
	if len(r.Referer()) != 0 && !strings.HasPrefix(r.Referer(), auth.getProtoHost(r.Proto, r.Host)) {
		return nil, errors.New("detect CSRF attack from " + r.Referer())
	}

	access_token_cookie, err_cookie := r.Cookie("access_token")
	if err_cookie != nil {
		return nil, err_cookie
	}

	token, err_token := jwt.Parse(access_token_cookie.Value, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(auth.ServerSecret), nil
	})
	if err_token != nil {
		return nil, err_token
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}

func (auth *TokenAuth) getProtoHost(proto, host string) string {
	if strings.HasPrefix(proto, "HTTPS") {
		return "https://" + host
	} else if strings.HasPrefix(proto, "HTTP") {
		return "http://" + host
	} else {
		return ""
	}
}
