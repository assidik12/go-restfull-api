package main

import (
	"log"

	"github.com/assidik12/go-restfull-api/config"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	router := config.InitializedServer()

	err := router.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server is starting on port 3000...")
}
