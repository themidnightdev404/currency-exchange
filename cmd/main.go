package main

import (
	"currency-exchange/internal/database"
	"currency-exchange/internal/handlers"     // Импортируем хендлеры
	"currency-exchange/internal/repositories" // Импортируем репозитории
	"fmt"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

func main() {
	// 1. Открываем БД
	db, err := database.OpenDB("database/exchange.db")
	if err != nil {
		log.Fatalf("Не удалось запустить БД: %v", err)
	}
	defer db.Close()
	fmt.Println("Успешно подключились к SQLite!")

	// 2. Запуск SQL-миграций
	migrationPath := "migrations/001_init.sql"
	err = database.RunMigrations(db, migrationPath)
	if err != nil {
		log.Fatalf("Ошибка применения миграции: %v", err)
	}
	fmt.Println("Миграция успешно применена!")

	// 3. СВЯЗЫВАЕМ СЛОИ (Dependency Injection)
	// Шаг А: Создаем репозиторий, отдаем ему базу
	repo := repositories.NewCurrencyRepository(db)
	// Шаг Б: Создаем хендлер, отдаем ему репозиторий
	currencyHandler := handlers.NewCurrencyHandler(repo)

	// 4. Настройка маршрутизатора
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", currencyHandler.PingHandler)
	mux.HandleFunc("/currencies", currencyHandler.GetCurrenciesHandler)

	// 5. Запуск сервера
	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
