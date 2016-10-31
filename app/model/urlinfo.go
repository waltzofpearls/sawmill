package model

import "time"

type UrlInfo struct {
	modelImpl
	Url         string
	Description string
	HasMalware  bool
	Created     time.Time
	Updated     time.Time
}

func NewUrlInfo(url, description string, hasMalware bool) *UrlInfo {
	mdl := &UrlInfo{
		Url:         url,
		Description: description,
		HasMalware:  hasMalware,
		Created:     time.Now(),
		Updated:     time.Now(),
	}
	mdl.SetId(url)
	return mdl
}

func (mdl *UrlInfo) GetId() string {
	return mdl.Url
}
