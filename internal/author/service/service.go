package service

import (
	"NewApiProd/internal/author/model"
	"NewApiProd/internal/author/storage"
	m1 "NewApiProd/internal/author/storage/model"
	"NewApiProd/pkg/api/filter"
	"NewApiProd/pkg/api/pagination"
	"NewApiProd/pkg/api/sort"
	"NewApiProd/pkg/logging"
	"context"
)

type Service struct {
	repository storage.Repository
	logger     *logging.Logger
}

func (s *Service) GetAll(ctx context.Context, ptoken pagination.POptions, filter filter.Options, sortOptions sort.Options) ([]model.Author, error) {
	options := m1.NewSortOptions(sortOptions.Field, sortOptions.Order)
	options1 := m1.NewFilterOptions(filter)
	options2 := m1.NewPaginationOptions(ptoken)
	all, err := s.repository.FindAll(ctx, options2, options1, options)
	if err != nil {
		s.logger.Fatalf("failed to get all authors due to error: %v", err)
	}

	return all, nil
}

func NewService(repository storage.Repository, logger *logging.Logger) *Service {
	return &Service{

		repository: repository,
		logger:     logger,
	}
}
