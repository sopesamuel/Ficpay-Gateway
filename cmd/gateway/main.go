package main

import (
	"database/sql"
	"flag"
	"os"
	"mart-gateway/internal/repository"
)

//configuration
//repository

func main() {

	dsn := flag.String("dsn", "gateway:gatewaypass@tcp(localhost:3306)/payment_gateway?parseTime=true", "MySQL data source name")
	db , err := openDB(*dsn)
	if err != nil {
		os.Exit(1)
	}
	defer db.Close()

	_ = repository.NewPaymentRepository(db)

}


func openDB(dsn string)(*sql.DB, error){

	db , err := sql.Open("mysql", dsn)
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
