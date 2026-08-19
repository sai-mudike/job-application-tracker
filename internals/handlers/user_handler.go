package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sai-mudike/job-application-tracker/internals/models"
	"github.com/sai-mudike/job-application-tracker/internals/repositories"
	"github.com/sai-mudike/job-application-tracker/internals/services"
)

func RegisterUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the user data provided"})
		return
	}

	err = repositories.CreateUser(&user)

	if err != nil {
		HandleError(context, err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})

}

func UserLogin(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, models.NewErrorResponse("INVALID_REQUEST", "Invalid request"))

		return
	}
	err = repositories.VerifyUser(&user)

	if err != nil {
		HandleError(context, err)
		return
	}

	token, err := services.GenerateToken(user.UserName, user.Id)

	if err != nil {
		HandleError(context, err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user succefully loged in", "token": token})

}
