package main

import (
	"github.com/gin-gonic/gin"
	"urlShortener/handlers"
	"urlShortener/service"
)

func multiplexer(router *gin.Engine, apiService *service.Service) {
	router.POST("shorten-url", handlers.CreateShortenUrl(apiService))
	router.GET(":key", handlers.GetOriginalUrl(apiService))
}
