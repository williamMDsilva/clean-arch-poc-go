package productusecase

import (
	"github.com/williamMDsilva/clean-arch-poc-go/internal/domain"
	"github.com/williamMDsilva/clean-arch-poc-go/internal/dto"
)

func (usecase usecase) Fetch(paginationRequest *dto.PaginationRequestParms) (*domain.Pagination[[]domain.Product], error) {
	products, err := usecase.repository.Fetch(paginationRequest)

	if err != nil {
		return nil, err
	}

	return products, nil
}
