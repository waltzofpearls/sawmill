package manager

import (
	"github.com/waltzofpearls/sawmill/app/model"
	"github.com/waltzofpearls/sawmill/app/repository"
)

type UrlInfo struct {
	urlInfoRepo repository.Repository
}

func NewUrlInfo(rpo *repository.UrlInfo) *UrlInfo {
	return &UrlInfo{
		urlInfoRepo: rpo,
	}
}

func (mgr *UrlInfo) GetUrlInfo(host, path, query string) (*model.UrlInfo, error) {
	url := host + "/" + path
	if query != "" {
		url += "?" + query
	}
	mdl, err := mgr.urlInfoRepo.Get(url, false)
	if err != nil {
		return nil, err
	}
	urlInfo := mdl.(*model.UrlInfo)
	return urlInfo, nil
}
