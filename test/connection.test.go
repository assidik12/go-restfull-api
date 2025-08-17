package test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func NewDB() (*sql.DB, error) {
	DB, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DATABASE"),
	))
	return DB, err
}
func TestNewDBConnection_Success_Ping(t *testing.T) {
	// Load .env file from the root directory for testing
	// Adjust the path based on your test execution directory
	err := godotenv.Load("../.env.test")
	if err != nil {
		log.Println("Warning: .env.test file not found, using environment variables")
	}

	// Ensure necessary env vars are set for the test database
	if os.Getenv("DB_USER") == "" || os.Getenv("DB_PASSWORD") == "" || os.Getenv("DB_HOST") == "" || os.Getenv("DB_PORT") == "" || os.Getenv("DB_NAME") == "" {
		t.Skip("Skipping database connection test: required environment variables are not set. Create a .env.test file.")
		return
	}

	db, err := NewDB()

	// Assert that no error occurred during connection
	assert.NoError(t, err)
	assert.NotNil(t, db)

	// Assert that the database is reachable
	err = db.Ping()
	assert.NoError(t, err)

	// Close the connection
	db.Close()
}

// TestNewDBConnection_Failure tests a failed database connection due to wrong credentials.
func TestNewDBConnection_Failure(t *testing.T) {
	// Set invalid database credentials temporarily
	t.Setenv("DB_HOST", "invalid-host")
	t.Setenv("DB_PORT", "0")
	t.Setenv("DB_USER", "invalid-user")
	t.Setenv("DB_PASSWORD", "invalid-password")
	t.Setenv("DB_NAME", "invalid-db")

	db, err := NewDB()

	// Assert that an error did occur
	assert.Error(t, err)
	assert.Nil(t, db)
}
