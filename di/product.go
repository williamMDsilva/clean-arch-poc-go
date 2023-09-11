package di

import (
	productservice "github.com/williamMDsilva/clean-arch-poc-go/adapter/api/product_service"
	"github.com/williamMDsilva/clean-arch-poc-go/adapter/postgres"
	productrepository "github.com/williamMDsilva/clean-arch-poc-go/adapter/postgres/product_repository"
	"github.com/williamMDsilva/clean-arch-poc-go/internal/domain"
	productusecase "github.com/williamMDsilva/clean-arch-poc-go/internal/domain/usecase/product_usecase"
)

// ConfigProductDI return a ProductService abstraction with dependency injection configuration
func ConfigProductDI(conn postgres.PoolInterface) domain.ProductService {
	productRepository := productrepository.New(conn)
	productUseCase := productusecase.New(productRepository)
	productService := productservice.New(productUseCase)

	return productService
}
