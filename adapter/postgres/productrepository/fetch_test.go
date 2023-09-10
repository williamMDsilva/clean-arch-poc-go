package productrepository_test

import (
	"fmt"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/require"
	"github.com/williamMDsilva/clean-arch-poc-go/adapter/postgres/productrepository"
	"github.com/williamMDsilva/clean-arch-poc-go/core/domain"
	"github.com/williamMDsilva/clean-arch-poc-go/core/dto"
)

func setupFetch() ([]string, dto.PaginationRequestParms, domain.Product, pgxmock.PgxPoolIface) {
	cols := []string{"id", "name", "price", "description"}
	fakePaginationRequestParams := dto.PaginationRequestParms{
		Page:         1,
		ItemsPerPage: 10,
		Sort:         nil,
		Descending:   nil,
		Search:       "",
	}
	fakeProductDBResponse := domain.Product{}
	faker.FakeData(&fakeProductDBResponse)

	mock, _ := pgxmock.NewPool()

	return cols, fakePaginationRequestParams, fakeProductDBResponse, mock
}

func TestFetch(t *testing.T) {
	cols, fakePaginationRequestParams, fakeProductDBResponse, mock := setupFetch()
	defer mock.Close()

	mock.ExpectQuery("SELECT (.+) FROM product").
		WillReturnRows(pgxmock.NewRows(cols).AddRow(
			fakeProductDBResponse.ID,
			fakeProductDBResponse.Name,
			fakeProductDBResponse.Price,
			fakeProductDBResponse.Description,
		))

	mock.ExpectQuery("SELECT COUNT(.+) FROM product").
		WillReturnRows(pgxmock.NewRows([]string{"count"}).AddRow(int32(1)))

	sut := productrepository.New(mock)
	products, err := sut.Fetch(&fakePaginationRequestParams)

	require.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// fix interface of products.Items
	for _, product := range products.Items {
		require.Nil(t, err)
		require.NotEmpty(t, product.ID)
		require.Equal(t, product.Name, fakeProductDBResponse.Name)
		require.Equal(t, product.Price, fakeProductDBResponse.Price)
		require.Equal(t, product.Description, fakeProductDBResponse.Description)
	}
}

func TestFetch_QueryError(t *testing.T) {
	_, fakePaginationRequestParams, _, mock := setupFetch()
	defer mock.Close()

	mock.ExpectQuery("SELECT (.+) FROM product").
		WillReturnError(fmt.Errorf("ANY QUERY ERROR"))

	sut := productrepository.New(mock)
	products, err := sut.Fetch(&fakePaginationRequestParams)

	require.NotNil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	require.Nil(t, products)
}

func TestFetch_QueryCountError(t *testing.T) {
	cols, fakePaginationRequestParams, fakeProductDBResponse, mock := setupFetch()
	defer mock.Close()

	mock.ExpectQuery("SELECT (.+) FROM product").
		WillReturnRows(pgxmock.NewRows(cols).AddRow(
			fakeProductDBResponse.ID,
			fakeProductDBResponse.Name,
			fakeProductDBResponse.Price,
			fakeProductDBResponse.Description,
		))

	mock.ExpectQuery("SELECT COUNT(.+) FROM product").
		WillReturnError(fmt.Errorf("ANY QUERY COUNT ERROR"))

	sut := productrepository.New(mock)
	products, err := sut.Fetch(&fakePaginationRequestParams)

	require.NotNil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	require.Nil(t, products)
}
