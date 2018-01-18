package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	// Connection Driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db  *gorm.DB
	err error
)

// Setup DB
func Setup() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	var connection string
	env := os.Getenv("ENVIROMENT_NAME")

	if env == "LOCAL" {
		connection = getLocalConnection()
	}

	if env == "" {
		connection = os.Getenv("CLEARDB_DATABASE_URL")
	}

	db, err = gorm.Open("mysql", connection)

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
}

// GetConnection return the connection to database
func GetConnection() *gorm.DB {
	return db
}

func getLocalConnection() string {

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
}
