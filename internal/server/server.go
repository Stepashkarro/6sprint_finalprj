package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// структура нашего сервера
type Server struct {
	logger *log.Logger
	http   *http.Server
}

// функция New создает, настраивает и возвращает новый сервер
func New(logger *log.Logger) *Server {
	// Создаем http-роутер
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	// настройки для сервера
	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		logger: logger,
		http:   httpServer,
	}
}

// запуск сервера и вывод в лог
func (s *Server) Run() error {
	s.logger.Println("Сервер запущен на порту :8080")
	return s.http.ListenAndServe()
}
