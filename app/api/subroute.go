package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/uber-go/zap"
	"github.com/waltzofpearls/sawmill/app/config"
	"github.com/waltzofpearls/sawmill/app/logger"
)

const (
	JsonContentTypeKey   = "Content-Type"
	JsonContentTypeValue = "application/json; charset=UTF-8"
)

type Subrouter interface {
	ConfigWith(*mux.Router, *config.Config, *logger.Logger)
	Handle()
}

type Subroute struct {
	Config *config.Config
	Logger *logger.Logger
	Router *mux.Router
}

func (sr *Subroute) ConfigWith(r *mux.Router, c *config.Config, l *logger.Logger) {
	sr.Router = r
	sr.Config = c
	sr.Logger = l
}

func (sr *Subroute) JsonResponseHandler(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.WriteHeader(http.StatusOK)
	sr.JsonBaseHandler(w, r, data)
}

func (sr *Subroute) JsonNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	jerr := JsonError{"Uh oh! You are requesting something that does not exist."}
	sr.JsonBaseHandler(w, r, jerr)
}

func (sr *Subroute) JsonInternalErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	jerr := JsonError{err.Error()}
	sr.JsonBaseHandler(w, r, jerr)
}

func (sr *Subroute) JsonBaseHandler(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set(JsonContentTypeKey, JsonContentTypeValue)
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
