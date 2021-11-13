package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"urlShortener/models"
)

func (service Service) CreateKey(c context.Context, url string) (string, error) {
	key, err := service.DaoService.QueryKey(c, models.ShortenedUrlMongoTableName(), url)
	if err == nil {
		return key, nil
	}
	key, err = service.RandomUniqueKeyGenrator.Generate()
	if err != nil {
		return "", err
	}
	shortenedUrl := models.ShortenedUrl{
		ID:        primitive.NewObjectID(),
		Key:       key,
		Url:       url,
		Length:    len(key),
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
		DeletedAt: time.Time{},
		IsDeleted: false,
	}
	collectionName := models.ShortenedUrlMongoTableName()
	service.DaoService.Save(c, collectionName, shortenedUrl)
	return key, err
}

func (service Service) GetUrl(c context.Context, key string) (string, error) {
	url, err := service.DaoService.GetUrl(c, models.ShortenedUrlMongoTableName(), key)
	if err != nil {
		return "", err
	}
	return url, nil
}
