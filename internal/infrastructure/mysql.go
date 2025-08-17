package infrastructure

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

func DatabaseConnection() *sql.DB {
	DB, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connecting to database...")

	err = DB.Ping()

	if err != nil {
		log.Fatal(err)
		panic(errors.New("connection to database failed"))
	}

	fmt.Println("connection to database success...")

	DB.SetMaxIdleConns(5)
	DB.SetMaxOpenConns(10)
	DB.SetConnMaxIdleTime(10 * time.Minute)
	DB.SetConnMaxLifetime(60 * time.Minute)

	return DB
}
