package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCatch404Handler(t *testing.T) {
	var (
		resp *httptest.ResponseRecorder
		req  *http.Request
	)

	r := new(mux.Router)
	r.Handle("/sup/bro", AttachMiddleware(
		r,
		Catch404Handler(),
		helpTestMiddleware(),
	))

	resp = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/sup/bro", nil)
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	resp = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/hey/dude", nil)
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestLoggingHandler(t *testing.T) {
	r := new(mux.Router)
	r.Handle("/bilbo/baggins", AttachMiddleware(
		r,
		LoggingHandler(ioutil.Discard),
		helpTestMiddleware(),
	))

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/bilbo/baggins", nil)
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func helpTestMiddleware() Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Yo man!")
		})
	}
}
