package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thebrodigy/sho.rt/handler"
)

func ApiRoutes(router *gin.Engine) {
	router.POST("shorten", handler.CreateShortUrl)
}
