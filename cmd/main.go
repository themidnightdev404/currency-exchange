package main

import (
	"currency-exchange/internal/database"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"currency-exchange/internal/repositories"

	_ "modernc.org/sqlite"
)

type application struct {
	db *sql.DB
}

func (app *application) pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func (app *application) getCurrenciesHandler(w http.ResponseWriter, r *http.Request) {
	repo := repositories.NewCurrencyRepository(app.db)
	list, err := repo.FindAllCurrencies()
	if err != nil {
		http.Error(w, "база данных недоступна", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func main() {
	// 1. Открываем БД
	db, err := database.OpenDB("database/exchange.db")
	if err != nil {
		log.Fatalf("Не удалось запустить БД: %v", err)
	}
	defer db.Close()
	fmt.Println("Успешно подключились к SQLite!")

	// 2. Запуск SQL-миграций для создания таблиц
	migrationPath := "migrations/001_init.sql"
	err = database.RunMigrations(db, migrationPath)
	if err != nil {
		log.Fatalf("Ошибка применения миграции: %v", err)
	}
	fmt.Println("Миграция успешно применена!")

	// 3. Настройка приложения и маршрутизатора
	app := &application{db: db}
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", app.pingHandler)
	mux.HandleFunc("/currencies", app.getCurrenciesHandler)

	// 4. Запуск сервера
	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
