package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/uber-go/zap"
	"github.com/waltzofpearls/sawmill/app/config"
	"github.com/waltzofpearls/sawmill/app/database"
	"github.com/waltzofpearls/sawmill/app/logger"
)

const (
	JsonContentTypeKey   = "Content-Type"
	JsonContentTypeValue = "application/json"
)

type Subrouter interface {
	ConfigWith(*mux.Router, *database.Database, *config.Config, *logger.Logger)
	Handle()
}

type Subroute struct {
	Config   *config.Config
	Database *database.Database
	Logger   *logger.Logger
	Router   *mux.Router
}

func (sr *Subroute) ConfigWith(r *mux.Router, d *database.Database, c *config.Config, l *logger.Logger) {
	sr.Router = r
	sr.Config = c
	sr.Logger = l
	sr.Database = d
}

func (sr *Subroute) JsonResponseHandler(w http.ResponseWriter, r *http.Request, data interface{}) {
	sr.JsonBaseHandler(w, r, http.StatusOK, data)
}

func (sr *Subroute) JsonNotFoundHandler(w http.ResponseWriter, r *http.Request, err error) {
	var msg string
	if err == nil {
		msg = "Uh oh! You are requesting something that does not exist."
	} else {
		msg = err.Error()
	}
	jerr := JsonError{http.StatusNotFound, msg}
	sr.JsonBaseHandler(w, r, http.StatusNotFound, jerr)
}

func (sr *Subroute) JsonInternalErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	jerr := JsonError{http.StatusInternalServerError, err.Error()}
	sr.JsonBaseHandler(w, r, http.StatusInternalServerError, jerr)
}

func (sr *Subroute) JsonBaseHandler(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set(JsonContentTypeKey, JsonContentTypeValue)
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		sr.Logger.Error(
			"Internal server error.",
			zap.Error(err),
		)
		http.Error(
			w,
			"Oops! Internal server error :(",
			http.StatusInternalServerError,
		)
	}
}
