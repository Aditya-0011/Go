package middlewares

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type MiddlewareOptions int8

const (
	None MiddlewareOptions = iota
	AuthRequired
	ValidationRequired
	AuthAndValidationRequired
)

func AddMiddlewares(h http.Handler, mws ...Middleware) http.Handler {
	for i := len(mws) - 1; i >= 0; i-- {
		h = mws[i](h)
	}
	return h
}
