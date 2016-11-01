package manager

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/waltzofpearls/sawmill/app/database"
	"github.com/waltzofpearls/sawmill/app/model"
	"github.com/waltzofpearls/sawmill/app/repository"
)

type FakeUrlInfoRepo struct{ repository.Repository }

func (rpo *FakeUrlInfoRepo) Get(key string, notFoundOk bool) (model.Model, error) {
	return &model.UrlInfo{Url: key}, nil
}
func (rpo *FakeUrlInfoRepo) Save(mdl model.Model) (model.Model, error) { return nil, nil }
func (rpo *FakeUrlInfoRepo) getBucketName() string                     { return "" }
func (rpo *FakeUrlInfoRepo) getModel() model.Model                     { return nil }
func (rpo *FakeUrlInfoRepo) getCluster() database.RiakCluster          { return nil }

var ErrUrlInfoRepoBadGet = errors.New("Something went wrong when fetching urlinfo.")

type FakeUrlInfoRepoBadGet struct{ *FakeUrlInfoRepo }

func (rpo *FakeUrlInfoRepoBadGet) Get(key string, notFoundOk bool) (model.Model, error) {
	return nil, ErrUrlInfoRepoBadGet
}

func TestGetUrlInfo(t *testing.T) {
	var (
		mdl *model.UrlInfo
		err error
	)

	mgr := &UrlInfo{}

	mgr.urlInfoRepo = &FakeUrlInfoRepoBadGet{}
	mdl, err = mgr.GetUrlInfo("host", "path", "query")
	assert.Error(t, err)
	assert.Nil(t, mdl)
	assert.EqualError(t, err, ErrUrlInfoRepoBadGet.Error())

	mgr.urlInfoRepo = &FakeUrlInfoRepo{}
	mdl, err = mgr.GetUrlInfo("host", "path", "query")
	assert.NoError(t, err)
	assert.NotNil(t, mdl)
	assert.Equal(t, "host/path?query", mdl.GetId())
}
