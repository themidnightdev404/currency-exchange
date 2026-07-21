package main

import (
	"currency-exchange/internal/database"
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "modernc.org/sqlite"
)

// application — это структура (пользовательский тип данных).
// Она нужна как "контейнер", чтобы объединить общие ресурсы приложения.
// В данном случае мы храним здесь указатель на пул соединений с базой данных (*sql.DB).
type application struct {
	db *sql.DB
}

func main() {
	// 1. Инициализация базы данных
	dbAddress := "database/exchange.db" // Путь к файлу базы данных SQLite

	// Вызываем твою функцию OpenDB. Она подготавливает пул соединений
	// и сразу проверяет связь с файлом через db.Ping().
	db, err := database.OpenDB(dbAddress)
	if err != nil {
		// Если база не открылась, пишем ошибку в консоль и аварийно завершаем программу.
		log.Fatal("Критическая ошибка инициализации БД:", err)
	}
	// defer гарантирует, что соединение с БД закроется автоматически,
	// когда функция main() закончит свою работу (при выключении сервера).
	defer db.Close()
	err = database.RunMigrations(db, "migrations/001_init.sql")
	if err != nil {
		log.Fatal("Ошибка применения миграций: ", err)
	}
	// 2. Создание экземпляра структуры (Dependency Injection)
	// Создаем объект app и передаем в его поле "db" наше открытое подключение.
	app := &application{db: db}

	// 3. Настройка маршрутизации (Роутинг)
	// Создаем ServeMux (мультиплексор). Он распределяет входящие HTTP-запросы по функциям.
	mux := http.NewServeMux()

	// Привязываем URL-путь "/ping" к методу pingHandler.
	// Метод принадлежит объекту app, поэтому внутри handler-а будет доступ к базе данных.
	mux.HandleFunc("/ping", app.pingHandler)

	// 4. Конфигурация HTTP-сервера
	// Явно настраиваем параметры сервера, чтобы он не зависал при плохом сетевом соединении.
	server := &http.Server{
		Addr:         ":8080",         // Порт, на котором сервер будет принимать запросы
		Handler:      mux,             // Передаем список наших маршрутов (мультиплексор)
		ReadTimeout:  5 * time.Second, // Максимальное время на чтение запроса от клиента
		WriteTimeout: 5 * time.Second, // Максимальное время на отправку ответа клиенту
	}

	log.Println("Сервер запущен на http://localhost:8080")

	// Запускаем бесконечный цикл прослушивания порта.
	// Если метод ListenAndServe завершится (например, из-за ошибки порта), log.Fatal остановит программу.
	log.Fatal(server.ListenAndServe())
}

// pingHandler — это метод структуры application.
// За счет того, что функция объявлена как метод (app *application),
// внутри нее можно писать app.db и делать любые запросы к базе данных.
func (app *application) pingHandler(w http.ResponseWriter, r *http.Request) {
	// Отправляем клиенту текстовый ответ "pong" и HTTP-статус 200 OK.
	w.Write([]byte("pong"))
}
