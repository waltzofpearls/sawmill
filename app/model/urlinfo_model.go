package model

import "time"

type UrlInfoModel struct {
	modelImpl
	Url         string
	Description string
	HasMalware  bool
	Created     time.Time
	Updated     time.Time
}

func NewUrlInfoModel(url, description string, hasMalware bool) *UrlInfoModel {
	um := &UrlInfoModel{
		Url:         url,
		Description: description,
		HasMalware:  hasMalware,
		Created:     time.Now(),
		Updated:     time.Now(),
	}
	um.SetId(url)
	return um
}

func (um *UrlInfoModel) GetId() string {
	return um.Url
}
