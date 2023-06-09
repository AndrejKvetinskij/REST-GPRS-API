package db

import (
	"NewApiProd/internal/user"
	"NewApiProd/pkg/client/postgresql"
	"NewApiProd/pkg/logging"
	"context"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

// Create implements user.Storage
func (r *repository) Create(ctx context.Context, user user.User) (string, error) {
	panic("unimplemented")
}

// Delete implements user.Storage
func (r *repository) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// FindAll implements user.Storage
func (r *repository) FindAll(ctx context.Context) (u []user.User, err error) {
	panic("unimplemented")
}

// FindOne implements user.Storage
func (r *repository) FindOne(ctx context.Context, id string) (user.User, error) {
	panic("unimplemented")
}

// Update implements user.Storage
func (r *repository) Update(ctx context.Context, user user.User) error {
	panic("unimplemented")
}

func NewRepository(client postgresql.Client, logger *logging.Logger) user.Storage {
	return &repository{
		client: client,
		logger: logger,
	}
}
