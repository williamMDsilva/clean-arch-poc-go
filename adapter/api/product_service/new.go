package productservice

import "github.com/williamMDsilva/clean-arch-poc-go/internal/domain"

type service struct {
	usecase domain.ProductUseCase
}

// New returns contract implementation of ProductService
func New(usecase domain.ProductUseCase) domain.ProductService {
	return &service{
		usecase: usecase,
	}
}
