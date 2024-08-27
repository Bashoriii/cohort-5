package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func DatabaseConnection() *sql.DB {
	psqlInfo := "postgres://postgres:admin@localhost/learn_db?sslmode=disable"

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// fmt.Println("Database connected!")
	return db
}
