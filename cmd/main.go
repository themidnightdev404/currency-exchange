package main

import (
	"currency-exchange/internal/database"
	"currency-exchange/internal/repositories"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

type application struct {
	db *sql.DB
}

func (app *application) pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func main() {
	// 1. Открываем БД
	db, err := database.OpenDB("database/exchange.db")
	if err != nil {
		log.Fatalf("Не удалось запустить БД: %v", err)
	}
	defer db.Close()
	fmt.Println("Успешно подключились к SQLite!")

	// 2. ВЫЗЫВАЕМ ВАШУ НОВУЮ ФУНКЦИЮ МИГРАЦИЙ
	// Передаем ей открытую базу и путь к файлу
	migrationPath := "migrations/001_init.sql"
	err = database.RunMigrations(db, migrationPath)
	if err != nil {
		log.Fatalf("Ошибка применения миграции: %v", err)
	}
	fmt.Println("Миграция успешно применена через database.RunMigrations!")

	// 3. Инициализируем репозиторий валют
	repo := repositories.NewCurrencyRepository(db)

	// 4. ВЫЗЫВАЕМ МЕТОД РЕПОЗИТОРИЯ И ВЫВОДИМ В КОНСОЛЬ
	list, err := repo.FindAllCurrencies()
	if err != nil {
		log.Fatalf("Не удалось получить валюты: %v", err)
	}

	fmt.Println("\n--- СПИСОК ВАЛЮТ ИЗ БАЗЫ ДАННЫХ ---")
	for _, c := range list {
		fmt.Printf("ID: %d | Код: %s | Название: %s | Знак: %s\n", c.ID, c.Code, c.FullName, c.Sign)
	}
	fmt.Println("----------------------------------\n")

	// 5. Настройка и запуск сервера
	mux := http.NewServeMux()
	app := &application{db: db}
	mux.HandleFunc("/ping", app.pingHandler)

	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
