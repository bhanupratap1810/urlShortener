package service

import (
	"github.com/teris-io/shortid"
	"urlShortener"
)

type Service struct {
	DaoService              urlShortener.Dao
	RandomUniqueKeyGenrator *shortid.Shortid
}
