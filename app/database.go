package app

import (
	"database/sql"
	"fmt"
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

	godotenv.Load(".env")
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	DB, err := sql.Open("mysql", dsn)

	helper.PanicError(err)
	DB.SetMaxIdleConns(5)
	DB.SetMaxOpenConns(10)
	DB.SetConnMaxIdleTime(10 * time.Minute)
	DB.SetConnMaxLifetime(60 * time.Minute)

	return DB
}
