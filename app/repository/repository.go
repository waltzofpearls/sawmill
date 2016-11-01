package repository

import (
	"encoding/json"
	"errors"

	riak "github.com/basho/riak-go-client"
	"github.com/waltzofpearls/sawmill/app/database"
	"github.com/waltzofpearls/sawmill/app/model"
)

var (
	ErrKeyNotFound        = errors.New("Cannot find the given key in the bucket.")
	ErrUnexpectedSiblings = errors.New("Unexpected siblings in response!")
)

type Repository interface {
	Get(key string, notFoundOk bool) (model.Model, error)
	Save(model.Model) (model.Model, error)
	getBucketName() string
	getModel() model.Model
	getCluster() database.RiakCluster
}

type repositoryImpl struct {
	cluster database.RiakCluster
}

func (r *repositoryImpl) getCluster() database.RiakCluster {
	return r.cluster
}

func get(r Repository, key string, notFoundOk bool) (model.Model, error) {
	cluster := r.getCluster()
	bucket := r.getBucketName()
	cmd, err := riak.NewFetchValueCommandBuilder().
		WithBucket(bucket).
		WithKey(key).
		WithNotFoundOk(notFoundOk).
		Build()
	if err != nil {
		return nil, err
	}
	if err = cluster.Execute(cmd); err != nil {
		return nil, err
	}

	fcmd := cmd.(*riak.FetchValueCommand)

	if len(fcmd.Response.Values) == 0 {
		if notFoundOk {
			return nil, nil
		}
		return nil, ErrKeyNotFound
	}

	if len(fcmd.Response.Values) > 1 {
		return nil, ErrUnexpectedSiblings
	} else {
		return buildModel(r.getModel(), fcmd.Response.Values[0])
	}
}

func save(r Repository, m model.Model) (model.Model, error) {
	cluster := r.getCluster()
	bucket := r.getBucketName()
	key := m.GetId()

	cmd, err := riak.NewFetchValueCommandBuilder().
		WithBucket(bucket).
		WithKey(key).
		WithNotFoundOk(true).
		Build()
	if err != nil {
		return nil, err
	}
	if err = cluster.Execute(cmd); err != nil {
		return nil, err
	}

	modelJson, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	var objToInsertOrUpdate *riak.Object
	fcmd := cmd.(*riak.FetchValueCommand)
	if len(fcmd.Response.Values) > 1 {
		objToInsertOrUpdate = fcmd.Response.Values[0]
		objToInsertOrUpdate.Value = modelJson
	} else {
		objToInsertOrUpdate = &riak.Object{
			Bucket:      bucket,
			Key:         key,
			ContentType: "application/json",
			Charset:     "utf8",
			Value:       modelJson,
		}
	}

	cmd, err = riak.NewStoreValueCommandBuilder().
		WithContent(objToInsertOrUpdate).
		WithReturnBody(true).
		Build()
	if err != nil {
		return nil, err
	}
	if err = cluster.Execute(cmd); err != nil {
		return nil, err
	}

	scmd := cmd.(*riak.StoreValueCommand)
	if len(scmd.Response.Values) > 1 {
		return nil, ErrUnexpectedSiblings
	}
	obj := scmd.Response.Values[0]
	return buildModel(r.getModel(), obj)
}

func buildModel(m model.Model, obj *riak.Object) (model.Model, error) {
	err := json.Unmarshal(obj.Value, m)
	m.SetId(obj.Key)
	return m, err
}
