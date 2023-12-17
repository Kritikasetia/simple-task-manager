// internal/db/database.go

package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var dbConn *sql.DB

func InitDB() {
	var err error
	//replace it with the actual mysql username and password
	dataSourceName := "<username>:<password>@tcp(localhost:3306)/taskmanagerdb"
	dbConn, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	err = dbConn.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to the database")
}

func GetDB() *sql.DB {
	return dbConn
}

func CloseDB() {
	dbConn.Close()
}
