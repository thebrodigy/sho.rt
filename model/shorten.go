package model

import "gorm.io/gorm"

type ShortenRequest struct {
	Url string `json:"url"`
}

type ShortenResponse struct {
	ShortUrl string `json:"shortUrl"`
}

type ShortUrl struct {
	gorm.Model
	OriginalUrl string `json:"originalUrl" gorm:"uniqueIndex"`
	ShortCode   string `json:"shortCode"`
}
