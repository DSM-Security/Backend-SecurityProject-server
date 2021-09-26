package db

import (
	"database/sql"
	"log"
)

var db *sql.DB
var err error

func Start() {
	db, err = sql.Open("mysql", "root:4451@tcp(127.0.0.1:3306)/security")
	if err != nil {
		log.Panic(err)
	}
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() {
	db.Close()
}
