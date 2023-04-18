package model

import (
	"NewApiProd/internal/author/storage"
	"NewApiProd/pkg/api/filter"
)

type filterOptions struct {
	limit  int
	fields []filter.Field
}

func NewFilterOptions(options filter.Options) storage.FilterOptions {
	return &filterOptions{
		limit:  options.Limit(),
		fields: options.Fields(),
	}
}

func (f *filterOptions) GetLimit() int {
	return f.limit
}

func (f *filterOptions) GetDateForQuery() []filter.Field {
	return f.fields
}
