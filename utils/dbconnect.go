package utils

import (
	"database/sql"
	"fmt"

	_ "gopkg.in/go-sql-driver/mysql.v1"
)

var db *sql.DB

// GetMongoDB function to return DB connection
func GetDBconn() *sql.DB {
	dbName := "AccountCrudDatabase"
	fmt.Println("conn info:", dbName)
	db, err := sql.Open("mysql",
		"root:root@tcp(mysql:3306)/AccountCrudDatabase")
	if err != nil {
		panic(err.Error())
	}
	// defer db.Close()

	return db
}
