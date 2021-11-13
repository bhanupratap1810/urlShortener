package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ShortenUrlRequestBody struct {
	Url string `json:"url"`
}

type ShortenedUrl struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Key       string             `json:"key" bson:"key"`
	Url       string             `json:"url" bson:"url"`
	Length    int                `json:"length" bson:"length"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	DeletedAt time.Time          `json:"deleted_at" bson:"deleted_at"`
	IsDeleted bool               `json:"is_deleted" bson:"is_deleted"`
}

func ShortenedUrlMongoTableName() string{
	return "shortened-url"
}