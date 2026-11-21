package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// логгер
	logger := log.New(os.Stdout, "APP: ", log.Ldate|log.Ltime|log.Lshortfile)

	// сервер
	srv := server.New(logger)

	// запуск и проверка на ошибку
	if err := srv.Run(); err != nil {
		logger.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
