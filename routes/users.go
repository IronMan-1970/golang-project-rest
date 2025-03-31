package routes

import (
	
	"go/by/example/restful/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func saveUser(context *gin.Context) {
	var user models.Users
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Binding failed", "error": err})
    return
	}
	
  err = user.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Save failed", "error": err})
    return
  }
  context.JSON(http.StatusCreated, gin.H{"message": "Save sucssed", "error": err})
}

func login(context *gin.Context) {}
