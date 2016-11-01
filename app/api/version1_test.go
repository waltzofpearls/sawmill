package api

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	riak "github.com/basho/riak-go-client"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/waltzofpearls/sawmill/app/database"
)

func TestNotFoundHandler(t *testing.T) {
	v1 := &Version1{}
	v1.Router = new(mux.Router)
	v1.Handle()

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	v1.Router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusNotFound, resp.Code)
}

type FakeRiakClusterFetcherNotFound struct{ database.RiakCluster }

func (c *FakeRiakClusterFetcherNotFound) Execute(cmd riak.Command) error {
	fcmd := cmd.(*riak.FetchValueCommand)
	fcmd.Response = &riak.FetchValueResponse{
		Values: []*riak.Object{},
	}
	return nil
}

var ErrRiakClusterFetcherError = errors.New("Uh oh!")

type FakeRiakClusterFetcherError struct{ database.RiakCluster }

func (c *FakeRiakClusterFetcherError) Execute(cmd riak.Command) error {
	return ErrRiakClusterFetcherError
}

type FakeRiakClusterFetcher struct{ database.RiakCluster }

func (c *FakeRiakClusterFetcher) Execute(cmd riak.Command) error {
	fcmd := cmd.(*riak.FetchValueCommand)
	fcmd.Response = &riak.FetchValueResponse{
		Values: []*riak.Object{
			&riak.Object{
				Value: []byte("{}"),
			},
		},
	}
	return nil
}

func TestUrlHandler(t *testing.T) {
	var (
		resp *httptest.ResponseRecorder
		req  *http.Request
	)

	db := &database.Database{}
	v1 := &Version1{}
	v1.Router = new(mux.Router)
	v1.Database = db
	v1.Handle()

	db.Cluster = &FakeRiakClusterFetcherNotFound{}
	resp = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/host/path?query=value", nil)
	v1.Router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusNotFound, resp.Code)

	db.Cluster = &FakeRiakClusterFetcherError{}
	resp = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/host/path?query=value", nil)
	v1.Router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusInternalServerError, resp.Code)

	db.Cluster = &FakeRiakClusterFetcher{}
	resp = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/host/path?query=value", nil)
	v1.Router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}
