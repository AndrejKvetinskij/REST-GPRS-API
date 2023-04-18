package storage

import (
	author "NewApiProd/internal/author/model"
	"NewApiProd/pkg/api/filter"
	"context"
)

type Repository interface {
	Create(ctx context.Context, author author.Author) error
	FindOne(ctx context.Context, id string) (author.Author, error)
	Update(ctx context.Context, author author.Author) error
	Delete(ctx context.Context, id string) error
	FindAll(ctx context.Context, paginationOptions PaginationOptions, filterOptions FilterOptions, sortOptions SortOptions) (a []author.Author, err error)
}

type SortOptions interface {
	GetOrderBy() string
}

type FilterOptions interface {
	GetLimit() int

	GetDateForQuery() []filter.Field
}

type PaginationOptions interface {
	GetPtoken() int
}
