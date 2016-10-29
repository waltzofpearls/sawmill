package api

import (
	"io"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Middleware func(http.Handler) http.Handler

func AttachMiddleware(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}

func Catch404Handler() Middleware {
	return func(h http.Handler) http.Handler {
		sr := &Subroute{}
		hf := http.HandlerFunc(sr.JsonNotFoundHandler)
		r, ok := h.(*mux.Router)
		if !ok {
			return hf
		}
		r.NotFoundHandler = hf
		return r
	}
}

func LoggingHandler(out io.Writer) Middleware {
	return func(h http.Handler) http.Handler {
		return handlers.CombinedLoggingHandler(out, h)
	}
}
