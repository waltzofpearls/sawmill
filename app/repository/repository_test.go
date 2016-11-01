package repository

import (
	"testing"

	riak "github.com/basho/riak-go-client"
	"github.com/stretchr/testify/assert"
	"github.com/waltzofpearls/sawmill/app/database"
	"github.com/waltzofpearls/sawmill/app/model"
)

type FakeRepository struct{ repositoryImpl }

func (r *FakeRepository) Get(key string, notFoundOk bool) (model.Model, error) { return nil, nil }
func (r *FakeRepository) Save(mdl model.Model) (model.Model, error)            { return nil, nil }

func (r *FakeRepository) getBucketName() string            { return "bucket" }
func (r *FakeRepository) getModel() model.Model            { return &model.UrlInfo{} }
func (r *FakeRepository) getCluster() database.RiakCluster { return r.cluster }

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

type FakeRiakClusterFetcherNotFound struct{ database.RiakCluster }

func (c *FakeRiakClusterFetcherNotFound) Execute(cmd riak.Command) error {
	fcmd := cmd.(*riak.FetchValueCommand)
	fcmd.Response = &riak.FetchValueResponse{
		Values: []*riak.Object{},
	}
	return nil
}

type FakeRiakClusterSaver struct{ database.RiakCluster }

func (c *FakeRiakClusterSaver) Execute(cmd riak.Command) error {
	switch cmd.(type) {
	case *riak.FetchValueCommand:
		fcmd := cmd.(*riak.FetchValueCommand)
		fcmd.Response = &riak.FetchValueResponse{
			Values: []*riak.Object{},
		}
	case *riak.StoreValueCommand:
		fcmd := cmd.(*riak.StoreValueCommand)
		fcmd.Response = &riak.StoreValueResponse{
			Values: []*riak.Object{
				&riak.Object{
					Value: []byte("{}"),
				},
			},
		}
	}
	return nil
}

func TestGet(t *testing.T) {
	var (
		clr database.RiakCluster
		mdl model.Model
		err error
	)

	rpo := &FakeRepository{}

	clr = &FakeRiakClusterFetcherNotFound{}
	rpo.cluster = clr
	mdl, err = get(rpo, "key", true)
	assert.NoError(t, err)
	assert.Nil(t, mdl)

	clr = &FakeRiakClusterFetcherNotFound{}
	rpo.cluster = clr
	mdl, err = get(rpo, "key", false)
	assert.Error(t, err)
	assert.Nil(t, mdl)
	assert.EqualError(t, err, ErrKeyNotFound.Error())

	clr = &FakeRiakClusterFetcher{}
	rpo.cluster = clr
	mdl, err = get(rpo, "key", false)
	assert.NoError(t, err)
	assert.NotNil(t, mdl)
}

func TestSave(t *testing.T) {
	var (
		clr     database.RiakCluster
		mdl     model.Model
		err     error
		urlInfo *model.UrlInfo
	)

	rpo := &FakeRepository{}

	clr = &FakeRiakClusterSaver{}
	rpo.cluster = clr
	urlInfo = &model.UrlInfo{Url: "url"}
	mdl, err = save(rpo, urlInfo)
	assert.NoError(t, err)
	assert.NotNil(t, mdl)
}
