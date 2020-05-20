package jwtutils

import (
	"testing"
	"time"
)

func TestNewAuth(t *testing.T) {
	auth := NewTokenAuth()
	t.Log(auth.TokenExpireIn)
	TokenExpireIn = 10 * time.Minute
	t.Log(auth.TokenExpireIn)
}

func TestDefault(t *testing.T) {
	auth := defaultAuth
	t.Log(auth.TokenExpireIn)
	TokenExpireIn = 30 * time.Minute
	t.Log(auth.TokenExpireIn)
}
