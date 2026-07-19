package main

import (
	"currency-exchange/internal/database"
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "modernc.org/sqlite"
)

type application struct {
	db *sql.DB
}

func main() {
	db, err := database.OpenDB("database/exchange.db")
	if err != nil {
		log.Fatal("Критическая ошибка инициализации БД:", err)
	}
	defer db.Close()

	app := &application{db: db}

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", app.pingHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}

func (app *application) pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
