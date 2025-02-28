package order_test

import (
	"bytes"
	"encoding/json"
	"godz/6-order-api-cart/internal/product"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"godz/6-order-api-cart/internal/order"
	"godz/6-order-api-cart/pkg/db"
)

func bootstrap() (*order.OrderHandler, sqlmock.Sqlmock, error) {
	base, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: base,
	}))
	if err != nil {
		return nil, nil, err
	}
	orderRepo := order.NewOrderRepo(&db.Db{
		DB: gormDB,
	})
	handler := order.OrderHandler{
		OrderRepo: orderRepo,
	}
	return &handler, mock, nil
}

func TestCreateOrderSuccess(t *testing.T) {
	handler, mock, err := bootstrap()
	if err != nil {
		t.Fatal(err)
		return
	}

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	orderRequest := order.OrderCreateRequest{
		UserID: "1",
		Products: []product.ProductCreateRequest{
			{Title: "testProduct", Description: "test", Price: 100, ImageURL: "testImage"},
		},
		ShippingAddress: "123 Street",
	}

	data, _ := json.Marshal(orderRequest)
	reader := bytes.NewReader(data)

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/orders", reader)
	if err != nil {
		t.Fatal(err)
		return
	}

	handler.CreateOrder()(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}
