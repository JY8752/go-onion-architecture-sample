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
	db, err := sql.Open("sqlite3", "./go_onion_architecture.db")
	if err != nil {
		log.Fatal(err)
	}

	// テーブル作成
	cmd := `CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name STRING,
		created_at DATETIME
	)`

	_, err = db.Exec(cmd)
	if err != nil {
		log.Fatal(err)
	}
}

func NewDBClient() *DBClient {
	return &DBClient{Client: db}
}
