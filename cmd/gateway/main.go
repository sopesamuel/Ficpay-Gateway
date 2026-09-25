package main

import (
	"database/sql"
	"flag"
	"net/http"
	"os"

	"mart-gateway/internal/api"
	"mart-gateway/internal/bank"
	"mart-gateway/internal/repository"
	"mart-gateway/internal/service"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dsn := flag.String("dsn", "gateway:gatewaypass@tcp(localhost:3306)/payment_gateway?parseTime=true", "MySQL data source name")
	bankURL := flag.String("bank-url", "http://localhost:8787", "Mock bank base URL")
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	db, err := openDB(*dsn)
	if err != nil {
		os.Exit(1)
	}
	defer db.Close()

	repo := repository.NewPaymentRepository(db)
	bankClient := bank.NewClient(*bankURL)
	mainService := service.ActivateService(bankClient, repo)
	handler := api.NewHandler(mainService)

	err = http.ListenAndServe(*addr, handler.Routes())
	if err != nil {
		os.Exit(1)
	}
}

func openDB(dsn string) (*sql.DB, error) {

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}