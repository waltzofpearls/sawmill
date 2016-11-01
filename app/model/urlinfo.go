package model

import "time"

type UrlInfo struct {
	modelImpl
	Url         string    `json:"url"`
	Description string    `json:"description"`
	HasMalware  bool      `json:"has_malware"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
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
