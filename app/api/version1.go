package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Version1 struct {
	Subroute
}

func (v1 *Version1) Handle() {
	v1.Router.HandleFunc("/", v1.notFoundHandler).Methods("GET")
	v1.Router.HandleFunc("/{host}/{path:.*}", v1.urlHandler).Methods("GET")
}

func (v1 *Version1) notFoundHandler(w http.ResponseWriter, r *http.Request) {
	v1.JsonNotFoundHandler(w, r)
}

func (v1 *Version1) urlHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vars["query"] = r.URL.RawQuery
	v1.JsonResponseHandler(w, r, vars)
}
