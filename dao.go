package urlShortener

import (
	"context"
	"urlShortener/models"
)

type Dao interface {
	Save(c context.Context, collectionName string, model models.ShortenedUrl) error
	QueryKey(c context.Context, collectionName string, url string) (string, error)
	GetUrl(c context.Context, collectionName string, key string) (string, error)
}
