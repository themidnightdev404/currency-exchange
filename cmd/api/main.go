package main

import (
	"log"

	"github.com/themidnightdev404/currency-exchange/database"
	"github.com/themidnightdev404/currency-exchange/internal/handlers"
)

func main() {

	port := "8080"

	db := database.InitDB()

	defer db.Close()

	h := handlers.NewCurrencyHandler()

	routes := h.InitRoutes()

	srv := handlers.NewServer(port, routes)

	if err := srv.Run(); err != nil {

		log.Fatalf("Критическая ошибка при работе HTTP-сервера: %v", err)
	}
}
