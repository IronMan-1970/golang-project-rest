package main

import (
	"go/by/example/restful/api/db"
	"go/by/example/restful/api/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	port := os.Getenv("PORT") // Отримання порту з середовища
	if port == "" {
		port = "8080" // Використовується для локального запуску
	}
	server.Run(":" + port) // Запускаємо сервер на отриманому порту
}
