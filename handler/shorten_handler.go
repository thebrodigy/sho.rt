package handler

import (
	"net/http"

	"errors"
	"math/rand"
	"os"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/thebrodigy/sho.rt/db"
	"github.com/thebrodigy/sho.rt/model"
)

var baseUrl string = os.Getenv("BASE_URL")

func CreateShortUrl(c *gin.Context) {
	var shortenRequest model.ShortenRequest

	if err := c.ShouldBindJSON(&shortenRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var shortUrl model.ShortUrl
	result := db.DB.Where("original_url = ?", shortenRequest.Url).First(&shortUrl)

	if result.Error == nil {
		c.JSON(http.StatusOK, shortUrl)
		return
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		shortUrl.OriginalUrl = shortenRequest.Url
		shortUrl.ShortCode = baseUrl + generateShortCode()

		if err := db.DB.Create(&shortUrl).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, shortUrl)
		return
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

}

func generateShortCode() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := make([]byte, 6)

	for i := range code {
		code[i] = chars[rand.Intn(len(chars))]
	}

	return string(code)
}

func Redirect(c *gin.Context) {
	shortCode := c.Param("shortCode")
	var shortUrl model.ShortUrl

	result := db.DB.Where("short_code = ?", shortCode).First(&shortUrl)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid short code"})
		return
	}

	c.Redirect(http.StatusFound, shortUrl.OriginalUrl)
}
