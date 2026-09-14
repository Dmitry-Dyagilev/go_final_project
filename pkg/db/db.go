package db

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

const schema string = `CREATE TABLE IF NOT EXISTS scheduler (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	date CHAR(8) NOT NULL DEFAULT "",
	title VARCHAR(64) NOT NULL DEFAULT "",
	comment TEXT NOT NULL DEFAULT "",
	repeat VARCHAR(128) NOT NULL DEFAULT ""
);
CREATE INDEX IF NOT EXISTS scheduler_date ON scheduler (date);`

var DB *sql.DB

func Init() error {
	exePath, err := os.Executable()
	if err != nil {
		return err
	}
	rootPath := filepath.Dir(exePath)
	dbPath := filepath.Join(rootPath, "scheduler.db")

	DB, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}
	_, err = DB.Exec(schema)
	if err != nil {
		return err
	}

	return nil
}
