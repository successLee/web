package web

import (
	"net/http"
)

type CookieConfig struct {
	*http.Cookie
}

func newCookieConfig() *CookieConfig {
	return &CookieConfig{&http.Cookie{Path: "/", MaxAge: 30 * 86400, HttpOnly: true}}
}
