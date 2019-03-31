package middleware

import (
	"net/http"
)

// Middleware related definitions
type Middleware interface {
	ServeNext(h http.Handler) http.Handler
}

func With(h http.Handler, ms ...Middleware) http.Handler {
	for _, m := range ms {
		h = m.ServeNext(h)
	}
	return h
}
