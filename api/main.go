package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"

	"github.com/andrefrco/gofin/api/handler"
	"github.com/andrefrco/gofin/api/middleware"
	"github.com/andrefrco/gofin/config"
	"github.com/andrefrco/gofin/repository"
	"github.com/andrefrco/gofin/usecase/transaction"
	"github.com/codegangsta/negroni"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func main() {
	var config = config.Get()
	db, err := sql.Open("postgres", config.DatabaseURL)
	if err != nil {
		log.Fatal(err.Error())
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database")
	}
	defer func() { _ = db.Close() }()

	transactionRepo := repository.NewTransaction(db)
	transactionService := transaction.NewService(transactionRepo)

	r := mux.NewRouter()
	//handlers
	n := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.NewLogger(),
	)
	//transaction
	handler.MakeTransactionHandlers(r, *n, transactionService)

	http.Handle("/", r)
	r.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + config.Port,
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
