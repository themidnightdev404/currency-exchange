package database

import (
	"database/sql"
	"fmt"
)

func OpenDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия: %w", err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("база недоступна: %w", err)
	}

	return db, nil
}
