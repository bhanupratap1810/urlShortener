package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"urlShortener/models"
	"urlShortener/service"
)

func CreateShortenUrl(s *service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody models.ShortenUrlRequestBody
		err := json.NewDecoder(c.Request.Body).Decode(&reqBody)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failure",
				"error":  "Error parsing request body",
			})
			return
		}
		key, err := s.CreateKey(c, reqBody.Url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "failure",
				"error":  "Error while creating the key",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"key":    key,
		})
	}
}

func GetOriginalUrl(s *service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var key string
		if key = c.Param("key"); key == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "failure",
				"message": "Key of shorten Url is missing",
			})
			return
		}
		url, err := s.GetUrl(c, key)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "failure",
				"message": "Error in querying the url",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"url":    url,
		})
	}
}
