package database

import (
	"database/sql"
	"fmt"
)

func OpenDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("проверка ping не удалась: %w", err)
	}

	return db, nil
}
