package jwtutils

import (
	"testing"
	"time"
)

func TestNewAuth(t *testing.T) {
	tokenAuth := NewTokenAuth()
	TokenExpireIn = 1 * time.Second

	if tokenAuth.TokenExpireIn == TokenExpireIn {
		t.Fatal("tokenAuth.TokenExpireIn should be the same")
	}
}

func TestDefault(t *testing.T) {
	auth := defaultAuth
	TokenExpireIn = 1 * time.Second

	if tokenAuth.TokenExpireIn != TokenExpireIn {
		t.Fatal("defaultAuth.TokenExpireIn should be changed")
	}
}
