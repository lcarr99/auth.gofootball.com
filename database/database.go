package database

import (
	"database/sql"
	"fmt"
	"sync"

	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var connection *sql.DB
var connectionError error
var once sync.Once

func NewDatabase() (*sql.DB, error) {
	once.Do(func() {
		godotenv.Load(".env")

		fmt.Println(os.Getenv("DB_HOST"))

		databaseConnection, openConnectionError := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_SCHEMA")))

		if openConnectionError != nil {
			connectionError = openConnectionError
			return
		}

		if databaseConnectionError := databaseConnection.Ping(); databaseConnectionError != nil {
			connectionError = databaseConnectionError
			return
		}

		connection = databaseConnection
	})

	return connection, connectionError
}
