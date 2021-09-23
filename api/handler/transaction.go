package handler

import (
	"encoding/json"
	"net/http"

	"github.com/andrefrco/gofin/usecase/transaction"

	"github.com/andrefrco/gofin/api/presenter"

	"github.com/andrefrco/gofin/entity"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func listTransactions(service transaction.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading transactions"
		var data []*entity.Transaction
		var err error
		title := r.URL.Query().Get("title")
		switch {
		case title == "":
			data, err = service.ListTransactions()
		default:
			data, err = service.SearchTransactions(title)
		}
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		var toJ []*presenter.Transaction
		for _, d := range data {
			toJ = append(toJ, &presenter.Transaction{
				ID:    d.ID,
				Title: d.Title,
				Value: d.Value,
			})
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func createTransaction(service transaction.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding transaction"
		var input struct {
			Title string `json:"title"`
			Value int64  `json:"value"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		id, err := service.CreateTransaction(input.Title, input.Value)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		toJ := &presenter.Transaction{
			ID:    id,
			Title: input.Title,
			Value: input.Value,
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func getTransaction(service transaction.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading transaction"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		data, err := service.GetTransaction(id)
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		toJ := &presenter.Transaction{
			ID:    data.ID,
			Title: data.Title,
			Value: data.Value,
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func deleteTransaction(service transaction.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error removing transactionmark"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		err = service.DeleteTransaction(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

//MakeTransactionHandlers make url handlers
func MakeTransactionHandlers(r *mux.Router, n negroni.Negroni, service transaction.UseCase) {
	r.Handle("/v1/transaction", n.With(
		negroni.Wrap(listTransactions(service)),
	)).Methods("GET", "OPTIONS").Name("listTransactions")

	r.Handle("/v1/transaction", n.With(
		negroni.Wrap(createTransaction(service)),
	)).Methods("POST", "OPTIONS").Name("createTransaction")

	r.Handle("/v1/transaction/{id}", n.With(
		negroni.Wrap(getTransaction(service)),
	)).Methods("GET", "OPTIONS").Name("getTransaction")

	r.Handle("/v1/transaction/{id}", n.With(
		negroni.Wrap(deleteTransaction(service)),
	)).Methods("DELETE", "OPTIONS").Name("deleteTransaction")
}
