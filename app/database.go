package app

import (
	"4kbr/restful-golang/helper"
	"database/sql"
	"os"
	"time"
)

func NewDB() *sql.DB {
	//user:password@tcp(localhost:5555)/dbname?tls=skip-verify&autocommit=true
	dbConnection := os.Getenv("DB_CONNECTION_STRING")
	db, err := sql.Open("mysql", dbConnection)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
