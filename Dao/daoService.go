package Dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"urlShortener/config"
	"urlShortener/models"
)

type DaoService struct {
	MongoConn *mongo.Database
}

func NewDaoService(ctx context.Context, mongoUri string) (*DaoService, error) {
	mongoConn, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err = mongoConn.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &DaoService{
		MongoConn: mongoConn.Database(config.MongoDbName),
	}, nil
}

func (s *DaoService) Save(c context.Context, collectionName string, model models.ShortenedUrl) error {
	col := s.MongoConn.Collection(collectionName)
	_, err := col.InsertOne(c, model)
	if err != nil {
		return err
	}
	return nil
}

func (s *DaoService) GetUrl(c context.Context, collectionName string, key string) (string, error) {
	col := s.MongoConn.Collection(collectionName)
	attr := models.ShortenedUrl{}
	resp := col.FindOne(c, bson.M{
		"key": key,
	})
	err := resp.Decode(&attr)
	if err != nil {
		return "", err
	}
	return attr.Url, nil
}

func (s *DaoService) QueryKey(c context.Context, collectionName string, url string) (string, error) {
	col := s.MongoConn.Collection(collectionName)
	attr := models.ShortenedUrl{}
	resp := col.FindOne(c, bson.M{
		"url": url,
	})
	err := resp.Decode(&attr)
	if err != nil {
		return "", err
	}
	return attr.Key, nil
}
