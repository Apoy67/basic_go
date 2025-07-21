package app

import (
	"basic-restfull-golang/halper"
	"database/sql"
	"time"
)

func NewDB() *sql.DB {
	dataSource := "root:rahasia123@tcp(localhost:3306)/dev_restfull?parseTime=true"
	db, err := sql.Open("mysql", dataSource)
	halper.PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
