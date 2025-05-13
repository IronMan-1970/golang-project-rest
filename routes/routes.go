package routes

import (
	"go/by/example/restful/api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authentification)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerToEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", saveUser)
	server.POST("/login", login)
}
