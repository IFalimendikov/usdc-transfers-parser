package storage 

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTable(db *sql.DB) error {
	// Create table if not exists
	createStmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS usdc_transfers (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			block_number INTEGER,
			sender TEXT,
			recipient TEXT,
			value TEXT
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}
	defer createStmt.Close()

	_, err = createStmt.Exec()
	if err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}
	return nil
}

func DropTable(db *sql.DB) error {
	// Drop table if exists
	dropStmt, err := db.Prepare(`DROP TABLE IF EXISTS usdc_transfers`)
	if err != nil {
		return fmt.Errorf("failed to prepare drop table statement: %v", err)
	}
	defer dropStmt.Close()

	_, err = dropStmt.Exec()
	if err != nil {
		return fmt.Errorf("failed to drop table: %v", err)
	}
	return nil
}
