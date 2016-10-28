package api

import "net/http"

type Version1 struct {
	Subroute
}

func (v1 *Version1) Handle() {
	v1.Router.Handle("/", Handler(v1.notFoundHandler)).Methods("GET")
	v1.Router.Handle("/{hostname_n_port}/{path_n_querystring}", Handler(v1.urlHandler)).Methods("GET")
}

func (v1 *Version1) notFoundHandler(w http.ResponseWriter, r *http.Request) error {
	return v1.JsonNotFoundHandler(w, r)
}

func (v1 *Version1) urlHandler(w http.ResponseWriter, r *http.Request) error {
	return v1.JsonResponseHandler(w, r, struct{}{})
}
