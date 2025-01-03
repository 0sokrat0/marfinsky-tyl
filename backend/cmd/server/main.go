package main

import (
	"backend/internal/repository"
	"backend/internal/routes"
	"log"
)

func main() {
	// Инициализация базы данных
	db, err := repository.InitDB("./data/fund.db")
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// Настройка маршрутов
	router := routes.SetupRoutes(db)
	router.SetTrustedProxies(nil)

	// Запуск сервера
	port := 8080
	log.Printf("Сервер запущен на http://localhost:%d", port)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
