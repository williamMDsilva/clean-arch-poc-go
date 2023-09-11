package productservice

import (
	"encoding/json"
	"net/http"

	"github.com/williamMDsilva/clean-arch-poc-go/internal/dto"
)

// @Summary Fetch products with server pagination
// @Description Fetch products with server pagination
// @Tags product
// @Accept  json
// @Produce  json
// @Param sort query string true "1,2"
// @Param descending query string true "true,false"
// @Param page query integer true "1"
// @Param itemsPerPage query integer true "10"
// @Param search query string false "some search string"
// @Success 200 "{object} domain.Pagination"
// @Router /product [get]
func (service service) Fetch(response http.ResponseWriter, request *http.Request) {
	paginationRequest, err := dto.FromValuePaginationRequestParams(request)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	products, err := service.usecase.Fetch(paginationRequest)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(products)
}
