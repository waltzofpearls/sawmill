package manager

import (
	"fmt"

	"github.com/waltzofpearls/sawmill/app/model"
	"github.com/waltzofpearls/sawmill/app/repository"
)

type UrlInfo struct {
	urlInfoRepo *repository.UrlInfo
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
	fmt.Println(url)
	mdl, err := mgr.urlInfoRepo.Get(url, false)
	if err != nil {
		return nil, err
	}
	urlInfo := mdl.(*model.UrlInfo)
	return urlInfo, nil
}
