package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"
)

func InitDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки файла .env: %v", err)
	}
	dbName := os.Getenv("DB_NAME")
	dbParams := os.Getenv("DB_PARAMS")
	dsn := fmt.Sprintf("%s%s", dbName, dbParams)
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		log.Fatalf("Не удалось открыть файл базы данных: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("База данных недоступна: %v", err)
	}
	return db
}
