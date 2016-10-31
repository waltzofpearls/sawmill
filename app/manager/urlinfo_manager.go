package manager

import (
	"fmt"

	"github.com/waltzofpearls/sawmill/app/model"
	"github.com/waltzofpearls/sawmill/app/repository"
)

type UrlInfoManager struct {
	urlInfoRepo *repository.UrlInfoRepository
}

func NewUrlInfoManager(ur *repository.UrlInfoRepository) *UrlInfoManager {
	return &UrlInfoManager{
		urlInfoRepo: ur,
	}
}

func (um *UrlInfoManager) GetUrlInfo(host, path, query string) (*model.UrlInfoModel, error) {
	url := host + "/" + path
	if query != "" {
		url += "?" + query
	}
	fmt.Println(url)
	m, err := um.urlInfoRepo.Get(url, false)
	if err != nil {
		return nil, err
	}
	urlInfo := m.(*model.UrlInfoModel)
	return urlInfo, nil
}
