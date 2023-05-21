package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DBClient struct {
	Client *sql.DB
}

var db *sql.DB

func init() {
	// sqlite3に接続
	var err error
	db, err = sql.Open("sqlite3", "./go_onion_architecture.db")
	if err != nil {
		log.Fatal(err)
	}

	// テーブル作成
	cmdList := []string{
		`CREATE TABLE IF NOT EXISTS users(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name STRING,
			created_at DATETIME
		)`,
		`CREATE TABLE IF NOT EXISTS todos(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER,
			title STRING,
			description STRING,
			created_at DATETIME,
			delete_at DATETIME,
			FOREIGN KEY(user_id) REFERENCES users(id)
		)`,
	}

	for _, cmd := range cmdList {
		_, err = db.Exec(cmd)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func NewDBClient() *DBClient {
	return &DBClient{Client: db}
}
