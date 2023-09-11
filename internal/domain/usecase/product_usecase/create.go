package productusecase

import (
	"github.com/williamMDsilva/clean-arch-poc-go/internal/domain"
	"github.com/williamMDsilva/clean-arch-poc-go/internal/dto"
)

func (usecase usecase) Create(productRequest *dto.CreateProductRequest) (*domain.Product, error) {
	product, err := usecase.repository.Create(productRequest)

	if err != nil {
		return nil, err
	}

	return product, nil
}
