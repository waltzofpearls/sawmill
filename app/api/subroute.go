package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/waltzofpearls/sawmill/app/config"
)

type Subrouter interface {
	ConfigWith(*mux.Router, *config.Config)
	Handle()
}

type Subroute struct {
	Router *mux.Router
	Config *config.Config
}

func (sr *Subroute) ConfigWith(r *mux.Router, c *config.Config) {
	sr.Router = r
	sr.Config = c
}

func (sr *Subroute) JsonResponseHandler(w http.ResponseWriter, r *http.Request, data interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		return sr.JsonErrorHandler(w, r, err)
	}
	return nil
}

func (sr *Subroute) JsonNotFoundHandler(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	return nil
}

func (sr *Subroute) JsonErrorHandler(w http.ResponseWriter, r *http.Request, err error) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(JsonError{err.Error()}); err != nil {
		panic(err)
	}
	return nil
}
