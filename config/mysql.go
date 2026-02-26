package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	dsn := "root@tcp(127.0.0.1:3306)/foobar"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("DB open error:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("DB connection error:", err)
	}

	log.Println("Connected to MySQL")

	return db
}