package routes

import (
	"go/by/example/restful/api/models"
	"go/by/example/restful/api/utils"
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

func login(context *gin.Context) {
	var user models.Users
	err := context.ShouldBind(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Binding failed", "error": err})
		return
	}
	err = user.Validate()
	if err != nil {
		context.JSON(
			http.StatusUnauthorized,
			gin.H{"message": "Authentification unsacsesful", "error": err},
		)
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Authentification unsacsesful", "error": err},
		)
		return
	}

  context.JSON(http.StatusOK, gin.H{"message": "Welcom", "token":token})

}
