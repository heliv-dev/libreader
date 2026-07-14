package myproject_sqlc

import (
	"context"
)

type Querier interface {
	DeleteBook(ctx context.Context, id int32) error
	GetBookAllData(ctx context.Context, id int32) (Books, error)
	InsertBook(ctx context.Context, arg InsertBookParams) (Books, error)
	ListBooks(ctx context.Context) ([]Books, error)
	ListBooksGenre(ctx context.Context, genre string) ([]Books, error)
	UpdateBook(ctx context.Context, arg UpdateBookParams) (Books, error)
	UpdateBookContent(ctx context.Context, arg UpdateBookContentParams) (Books, error)
}

var _ Querier = (*Queries)(nil)
