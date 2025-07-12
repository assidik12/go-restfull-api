package main

import (
	"log"

	"github.com/assidik12/go-restfull-api/config"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	server := config.InitializedServer()

	// PERUBAHAN: Pindahkan pesan log ke SEBELUM server dijalankan.
	// Gunakan log.Println agar formatnya konsisten dengan log Docker lainnya.
	log.Println("Server is starting on port 3000...")

	// ListenAndServe adalah fungsi blocking, ia akan berjalan selamanya di sini.
	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
