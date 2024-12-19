package main

import (
	"fmt"

	"github.com/assidik12/go-restfull-api/config"
	"github.com/assidik12/go-restfull-api/helper"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	server := config.InitializedServer()

	err := server.ListenAndServe()

	fmt.Println("Server running on port 8080")

	helper.PanicError(err)

}
