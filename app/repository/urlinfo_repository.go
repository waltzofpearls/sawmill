package repository

import (
	"time"

	riak "github.com/basho/riak-go-client"
	"github.com/waltzofpearls/sawmill/app/model"
)

type UrlInfoRepository struct {
	repositoryImpl
}

func NewUrlInfoRepository(c *riak.Cluster) *UrlInfoRepository {
	ur := &UrlInfoRepository{}
	ur.cluster = c
	return ur
}

func (ur *UrlInfoRepository) Get(key string, notFoundOk bool) (model.Model, error) {
	return get(ur, key, notFoundOk)
}

func (ur *UrlInfoRepository) Save(m model.Model) (model.Model, error) {
	um := m.(*model.UrlInfoModel)
	um.Updated = time.Now()
	return save(ur, um)
}

func (ur *UrlInfoRepository) getBucketName() string {
	return "urlinfo"
}

func (ur *UrlInfoRepository) getModel() model.Model {
	return &model.UrlInfoModel{}
}
