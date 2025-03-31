package main

import (
	"github.com/gin-gonic/gin"
	"go/by/example/restful/api/db"
	"go/by/example/restful/api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080") // localhost:8080 (до речі відки беруться порти у сайтах)
}
