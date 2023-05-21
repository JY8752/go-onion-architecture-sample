package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	createUsersTable = `
		CREATE TABLE IF NOT EXISTS users(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name STRING,
			created_at DATETIME
		)
	`
	createTodosTable = `
		CREATE TABLE IF NOT EXISTS todos(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER,
			title STRING,
			description STRING,
			created_at DATETIME,
			delete_at DATETIME,
			FOREIGN KEY(user_id) REFERENCES users(id)
		)
	`
)

type DBClient struct {
	Client *sql.DB
}

func createTables(db *sql.DB) error {
	cmdList := []string{createUsersTable, createTodosTable}

	for _, cmd := range cmdList {
		_, err := db.Exec(cmd)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewDBClient(datasource string) *DBClient {
	// sqlite3に接続
	db, err := sql.Open("sqlite3", datasource)
	if err != nil {
		log.Fatal(err)
	}

	// テーブル作成
	if err := createTables(db); err != nil {
		log.Fatal(err)
	}

	return &DBClient{Client: db}
}
