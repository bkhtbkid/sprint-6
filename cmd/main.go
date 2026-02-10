package main

import (
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.Default()
	srv := server.LoadServer(logger)
	err := srv.HTTPServer.ListenAndServe()
	if err != nil {
		srv.Logger.Fatal("Ошибка. Не удалось запустить сервер")
	}
}
