package di

import (
	"github.com/williamMDsilva/clean-arch-poc-go/adapter/http/productservice"
	"github.com/williamMDsilva/clean-arch-poc-go/adapter/postgres"
	"github.com/williamMDsilva/clean-arch-poc-go/adapter/postgres/productrepository"
	"github.com/williamMDsilva/clean-arch-poc-go/core/domain"
	"github.com/williamMDsilva/clean-arch-poc-go/core/domain/usecase/productusecase"
)

// ConfigProductDI return a ProductService abstraction with dependency injection configuration
func ConfigProductDI(conn postgres.PoolInterface) domain.ProductService {
	productRepository := productrepository.New(conn)
	productUseCase := productusecase.New(productRepository)
	productService := productservice.New(productUseCase)

	return productService
}
