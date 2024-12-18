package repositories

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "../../app.db")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(db)

	return db
}
