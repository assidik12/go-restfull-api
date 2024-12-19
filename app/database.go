package app

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/assidik12/go-restfull-api/helper"
	"github.com/joho/godotenv"
)

func SetupTestDB() *sql.DB {
	DBTEST, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go_rest_api?charset=utf8mb4&parseTime=True&loc=Local")

	helper.PanicError(err)
	DBTEST.SetMaxIdleConns(5)
	DBTEST.SetMaxOpenConns(10)
	DBTEST.SetConnMaxIdleTime(10 * time.Minute)
	DBTEST.SetConnMaxLifetime(60 * time.Minute)

	return DBTEST
}

func NewDB() *sql.DB {

	if godotenv.Load() != nil {
		log.Fatal("Error loading .env file")
	}

	dbName := os.Getenv("DB_NAME")

	DBPROD, err := sql.Open("mysql", "root:@tcp(localhost:3306)/"+dbName+"?charset=utf8mb4&parseTime=True&loc=Local")

	helper.PanicError(err)
	DBPROD.SetMaxIdleConns(5)
	DBPROD.SetMaxOpenConns(10)
	DBPROD.SetConnMaxIdleTime(10 * time.Minute)
	DBPROD.SetConnMaxLifetime(60 * time.Minute)

	return DBPROD
}
