package api

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestJsonResponseHandler(t *testing.T) {
	sr := &Subroute{}
	r := new(mux.Router)
	r.HandleFunc("/elrond/half-elven", func(w http.ResponseWriter, r *http.Request) {
		sr.JsonResponseHandler(w, r, map[string]string{
			"Title":   "Lord of Rivendell",
			"Species": "Elf, Half-elven",
			"Powers":  "Telepathy",
		})
	})

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/elrond/half-elven", nil)
	r.ServeHTTP(resp, req)

	expected := `{
		"Powers":"Telepathy",
		"Species":"Elf, Half-elven",
		"Title":"Lord of Rivendell"
	}`

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.JSONEq(t, expected, resp.Body.String())
}

func TestJsonNotFoundHandler(t *testing.T) {
	sr := &Subroute{}
	r := new(mux.Router)
	r.HandleFunc("/gandalf", func(w http.ResponseWriter, r *http.Request) {
		sr.JsonNotFoundHandler(w, r)
	})

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/gandalf", nil)
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestJsonInternalErrorHandler(t *testing.T) {
	err := errors.New("One ring to rule them all.")
	sr := &Subroute{}
	r := new(mux.Router)
	r.HandleFunc("/sauron", func(w http.ResponseWriter, r *http.Request) {
		sr.JsonInternalErrorHandler(w, r, err)
	})

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/sauron", nil)
	r.ServeHTTP(resp, req)

	expected := `{"errorMessage":"` + err.Error() + `"}`

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
	assert.JSONEq(t, expected, resp.Body.String())
}
