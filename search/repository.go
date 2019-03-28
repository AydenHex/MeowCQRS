package search

import (
	"context"
	"cqrs/schema"
)

type Repository interface {
	Close()
	InsertMeow(ctx context.Context, meow schema.Meow) error
	SearchMeow(ctx context.Context, query string, skip uint64, take uint64) ([]schema.Meow, error)
}

var impl Repository

func SetRepository(repository Repository) {

}
