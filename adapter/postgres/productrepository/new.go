package productrepository

import (
	"github.com/williamMDsilva/clean-arch-poc-go/adapter/postgres"
	"github.com/williamMDsilva/clean-arch-poc-go/core/domain"
)

type repository struct {
	db postgres.PoolInterface
}

// New returns contract implementation of ProductRepository
func New(db postgres.PoolInterface) domain.ProductRepository {
	return &repository{
		db: db,
	}
}
