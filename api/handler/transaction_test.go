package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/andrefrco/gofin/entity"

	"github.com/andrefrco/gofin/usecase/transaction/mock"
	"github.com/codegangsta/negroni"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Test_listTransactions(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeTransactionHandlers(r, *n, service)
	path, err := r.GetRoute("listTransactions").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/transaction", path)
	b := &entity.Transaction{
		ID: entity.NewID(),
	}
	service.EXPECT().
		ListTransactions().
		Return([]*entity.Transaction{b}, nil)
	ts := httptest.NewServer(listTransactions(service))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_createTransaction(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeTransactionHandlers(r, *n, service)
	path, err := r.GetRoute("createTransaction").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/transaction", path)

	service.EXPECT().
		CreateTransaction(gomock.Any(), gomock.Any()).
		Return(entity.NewID(), nil)
	h := createTransaction(service)

	ts := httptest.NewServer(h)
	defer ts.Close()
	payload := fmt.Sprintf(`{
		"title": "Supermarket",
		"value":100
	}`)
	resp, _ := http.Post(ts.URL+"/v1/transaction", "application/json", strings.NewReader(payload))
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var b *entity.Transaction
	json.NewDecoder(resp.Body).Decode(&b)
	assert.Equal(t, "Supermarket", b.Title)
}

func Test_listTransactions_NotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	ts := httptest.NewServer(listTransactions(service))
	defer ts.Close()
	service.EXPECT().
		SearchTransactions("transaction of transactions").
		Return(nil, entity.ErrNotFound)
	res, err := http.Get(ts.URL + "?title=transaction+of+transactions")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}

func Test_listTransactions_Search(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	b := &entity.Transaction{
		ID: entity.NewID(),
	}
	service.EXPECT().
		SearchTransactions("supermarket").
		Return([]*entity.Transaction{b}, nil)
	ts := httptest.NewServer(listTransactions(service))
	defer ts.Close()
	res, err := http.Get(ts.URL + "?title=supermarket")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_getTransaction(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeTransactionHandlers(r, *n, service)
	path, err := r.GetRoute("getTransaction").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/transaction/{id}", path)
	b := &entity.Transaction{
		ID: entity.NewID(),
	}
	service.EXPECT().
		GetTransaction(b.ID).
		Return(b, nil)
	handler := getTransaction(service)
	r.Handle("/v1/transaction/{id}", handler)
	ts := httptest.NewServer(r)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/v1/transaction/" + b.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	var d *entity.Transaction
	json.NewDecoder(res.Body).Decode(&d)
	assert.NotNil(t, d)
	assert.Equal(t, b.ID, d.ID)
}

func Test_deleteTransaction(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeTransactionHandlers(r, *n, service)
	path, err := r.GetRoute("deleteTransaction").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/transaction/{id}", path)
	b := &entity.Transaction{
		ID: entity.NewID(),
	}
	service.EXPECT().DeleteTransaction(b.ID).Return(nil)
	handler := deleteTransaction(service)
	req, _ := http.NewRequest("DELETE", "/v1/transaction/"+b.ID.String(), nil)
	r.Handle("/v1/transactionmark/{id}", handler).Methods("DELETE", "OPTIONS")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
