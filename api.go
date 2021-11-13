package urlShortener

import "context"

type ApiService interface {
	CreateKey(c context.Context, url string) (string, error)
	GetUrl(c context.Context, key string) (string, error)
}
