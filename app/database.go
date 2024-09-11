package app

import (
	"database/sql"
	"time"

	"github.com/assidik12/go-restfull-api/helper"
)

func SetupTestDB() *sql.DB {
	DBTEST, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go_rest_api_testing?charset=utf8mb4&parseTime=True&loc=Local")

	helper.PanicError(err)
	DBTEST.SetMaxIdleConns(5)
	DBTEST.SetMaxOpenConns(10)
	DBTEST.SetConnMaxIdleTime(10 * time.Minute)
	DBTEST.SetConnMaxLifetime(60 * time.Minute)

	return DBTEST
}

func NewDB() *sql.DB {
	DBPROD, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go_rest_api?charset=utf8mb4&parseTime=True&loc=Local")

	helper.PanicError(err)
	DBPROD.SetMaxIdleConns(5)
	DBPROD.SetMaxOpenConns(10)
	DBPROD.SetConnMaxIdleTime(10 * time.Minute)
	DBPROD.SetConnMaxLifetime(60 * time.Minute)

	return DBPROD
}
