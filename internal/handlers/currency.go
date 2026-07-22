package handlers

import (
	"currency-exchange/internal/repositories"
	"encoding/json"
	"log"
	"net/http"
)

// CurrencyHandler — это структура-контроллер (обработчик).
// Она хранит внутри себя указатель на репозиторий валют.
// Это поле ('repo') пишется с маленькой буквы, чтобы скрыть его от внешних пакетов (инкапсуляция).
type CurrencyHandler struct {
	repo *repositories.CurrencyRepository
}

// NewCurrencyHandler — это функция-конструктор.
// Она принимает уже готовый и подключенный к базе репозиторий, закладывает его
// внутрь структуры CurrencyHandler и возвращает указатель на этот созданный обработчик.
// Это гарантирует, что HTTP-методы никогда не останутся без доступа к данным.
func NewCurrencyHandler(repo *repositories.CurrencyRepository) *CurrencyHandler {
	return &CurrencyHandler{repo: repo}
}

// GetCurrenciesHandler — это метод структуры CurrencyHandler, который обрабатывает HTTP-запрос.
// Сигнатура (w, r) стандартна для всех веб-ручек в языке Go.
func (h *CurrencyHandler) GetCurrenciesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	// 1. ПОЛУЧЕНИЕ ДАННЫХ ИЗ БАЗЫ
	// Обработчик не умеет сам писать SQL-запросы. Вместо этого он обращается
	// к заложенному в него репозиторию (h.repo) и просит его вернуть список всех валют.
	list, err := h.repo.FindAllCurrencies()
	if err != nil {
		// 2. ОБРАБОТКА КРИТИЧЕСКОЙ ОШИБКИ БАЗЫ ДАННЫХ
		// Если репозиторий вернул ошибку, мы используем встроенную функцию http.Error.
		// Она делает две вещи: прерывает обычный ответ, отправляет клиенту статус 500 (Internal Server Error)
		http.Error(w, "база данных недоступна", http.StatusInternalServerError)
		return
	}

	// 3. НАСТРОЙКА HTTP-ЗАГОЛОВКОВ (HEADERS)
	// До того, как отправить сами данные, мы обязаны сказать браузеру или API-клиенту,
	// какой именно формат информации сейчас прилетит по сети.
	// Заголовок "Content-Type" со значением "application/json" сообщает, что сервер отдает JSON-текст.
	w.Header().Set("Content-Type", "application/json")

	// 4. КОДИРОВАНИЕ В JSON И ОТПРАВКА КЛИЕНТУ
	// json.NewEncoder(w) создает специальный инструмент кодирования, привязанный к сетевому потоку ответа (w).
	// Метод .Encode(list) берет наш срез Go-структур 'list', на лету превращает его
	// в текстовую строку формата JSON (с учетом наших json-тегов) и сразу же отправляет в сеть клиенту.
	err = json.NewEncoder(w).Encode(list)
	if err != nil {
		log.Printf("Ошибка кодирования JSON или отправки ответа: %v", err)
		return
	}
}
