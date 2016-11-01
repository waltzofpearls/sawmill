package repository

import (
	"time"

	"github.com/waltzofpearls/sawmill/app/database"
	"github.com/waltzofpearls/sawmill/app/model"
)

type UrlInfo struct {
	repositoryImpl
}

func NewUrlInfo(c database.RiakCluster) *UrlInfo {
	rpo := &UrlInfo{}
	rpo.cluster = c
	return rpo
}

func (rpo *UrlInfo) Get(key string, notFoundOk bool) (model.Model, error) {
	return get(rpo, key, notFoundOk)
}

func (rpo *UrlInfo) Save(mdl model.Model) (model.Model, error) {
	urlInfo := mdl.(*model.UrlInfo)
	urlInfo.Updated = time.Now()
	return save(rpo, urlInfo)
}

func (rpo *UrlInfo) getBucketName() string {
	return "urlinfo"
}

func (rpo *UrlInfo) getModel() model.Model {
	return &model.UrlInfo{}
}
