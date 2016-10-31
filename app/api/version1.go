package api

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/waltzofpearls/sawmill/app/manager"
	"github.com/waltzofpearls/sawmill/app/repository"
)

type Version1 struct {
	Subroute
}

func (v1 *Version1) Handle() {
	v1.Router.HandleFunc("/", v1.notFoundHandler).Methods("GET")
	v1.Router.HandleFunc("/{host}/{path:.*}", v1.urlHandler).Methods("GET")
}

func (v1 *Version1) notFoundHandler(w http.ResponseWriter, r *http.Request) {
	v1.JsonNotFoundHandler(w, r, nil)
}

func (v1 *Version1) urlHandler(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	c := v1.Database.Cluster

	rpo := repository.NewUrlInfoRepository(c)
	man := manager.NewUrlInfoManager(rpo)
	urlInfo, err := man.GetUrlInfo(v["host"], v["path"], r.URL.RawQuery)

	if err == repository.ErrKeyNotFound {
		v1.JsonNotFoundHandler(w, r, errors.New("Requested URL does not exist in the database."))
	} else if err != nil {
		v1.JsonInternalErrorHandler(w, r, err)
	} else {
		v1.JsonResponseHandler(w, r, urlInfo)
	}
}
