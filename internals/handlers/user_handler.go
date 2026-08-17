package handlers

import (
	"fmt"
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
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the user data provided"})
		return
	}

	err = repositories.CreateUser(&user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create the user"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "user created successfully"})

}

func UserLogin(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the user data provided"})
		return
	}

	err = repositories.VerifyUser(&user)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := services.GenerateToken(user.UserName, user.Id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create user token"})
	}

	context.JSON(http.StatusAccepted, gin.H{"message": "user succefully loged in", "token": token})

}
