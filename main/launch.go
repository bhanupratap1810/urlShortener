package main

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
	"log"
	"urlShortener/Dao"
	"urlShortener/config"
	"urlShortener/service"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	router := gin.New()

	router.Use(gin.Logger(), gin.Recovery())

	allowedHeaders := []string{config.ORIGIN, config.ACCEPT, config.ContentTypeHeader, config.AUTHORIZATION, config.DateUsed, config.XRequestedWith}

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "OPTIONS", "HEAD", "PUT", "DELETE"},
		AllowHeaders: allowedHeaders,
	}))

	auth := config.MongoUser + ":" + config.MongoPassword + "@"
	if auth == ":@" {
		auth = ""
	}
	mongoUri := "mongodb://" + auth + config.MongoServer
	daoSvc, err := Dao.NewDaoService(ctx, mongoUri)
	if err != nil {
		log.Fatal("unable to connect to mongoDb")
		return
	}

	sid, err := shortid.New(1, shortid.DefaultABC, 2342)

	apiService := &service.Service{
		DaoService: daoSvc,
		RandomUniqueKeyGenrator: sid,
	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	multiplexer(router, apiService)
	log.Print("Application loaded successfully ")
	log.Fatal(router.Run(":" + config.PORT))
}
