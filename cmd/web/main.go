package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/assidik12/go-restfull-api/config"
	"github.com/assidik12/go-restfull-api/helper"
)

func main() {
	server := config.InitializedServer()

	if godotenv.Load() != nil {
		log.Fatal("Error loading .env file")
	}

	err := server.ListenAndServe()

	helper.PanicError(err)

}
