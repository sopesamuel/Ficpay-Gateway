package main

import (
	"mart-gateway/internal/bank"
	"database/sql"
	"flag"
	"os"
)

//configuration
//repository

func main() {
	bank.SendRefundRequestToBank()

	dsn := flag.String("dsn", "web:pass@tcp(localhost:3306)/snippetbox?parseTime=true", "MySQL data source name") 

	_, err := openDB(*dsn)
	if err != nil {
		os.Exit(1)
	}
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
