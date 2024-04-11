package main

import (
	"fmt"
	"log"
	
	//sql
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	)

//package-level-variables

/**
  type DB struct{ ... }
    func Open(driverName, dataSourceName string) (*DB, error)
    func OpenDB(c driver.Connector) *DB
  */
var db *sql.DB

func	main() {
	//variables
	var dsn string = "root:1q2w3e4r@tcp(localhost:3306)/recordings"
	var err error

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Check connection between DB
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database")
}
