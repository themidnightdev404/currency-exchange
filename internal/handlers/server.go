package handlers

import (
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(port string, handler http.Handler) *Server {
	// return &Server{...} создает ваш объект в памяти, заполняет внутренний http.Server настройками и возвращает ссылку на него.
	return &Server{

		httpServer: &http.Server{
			Addr: ":" + port,

			Handler: handler,

			ReadTimeout: 6 * time.Second,

			WriteTimeout: 6 * time.Second,
		},
	}
}

// Run — метод структуры Server, который переводит сервер из режима «настройки» в режим «работы».
// СВЯЗАН С: Вызывается в самом конце функции main() в файле main.go.
// ЗАЧЕМ: Запускает бесконечный цикл прослушивания сети и обработки трафика.
func (s *Server) Run() error {
	// fmt.Printf выводит информационное сообщение в терминал (например, вашей Fedora).
	// СВЯЗАН С: Переменной s.httpServer.Addr, откуда берется настроенный порт.
	// ЗАЧЕМ: Чтобы разработчик визуально видел, что запуск начался и на каком адресе тестировать API.
	fmt.Printf("Сервер запускается на http://localhost%s\n", s.httpServer.Addr)

	// s.httpServer.ListenAndServe() — это ключевой метод, который захватывает порт в ОС и начинает принимать HTTP-запросы.
	// СВЯЗАН С: Ядром Linux/Файрволом (блокирует порт под ваше приложение).
	// ВАЖНО: Этот метод является блокирующим. Код «замирает» на этой строчке и крутится бесконечно.
	// Он вернет управление (и ошибку error) только если порт уже занят или сервер экстренно остановлен.
	return s.httpServer.ListenAndServe()
}
