package db

import (
	"database/sql"
	"log"
)

var db *sql.DB
var err error

func Start() {
	db, err = sql.Open("mysql", "root:4451@tcp(127.0.0.1:3306)/security")
	err = Migrate()
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

func Migrate() error {
	_, err := db.Query("CREATE TABLE user (id VARCHAR(20) PRIMARY KEY, password VARCHAR(255), nickname VARCHAR(30))")
	_, err = db.Query("CREATE TABLE post (pid INT(10) NOT NULL AUTO_INCREMENT PRIMARY KEY, title VARCHAR(20), content VARCHAR(100), writer VARCHAR(100), createdAt TIMESTAMP)")
	return err
}
