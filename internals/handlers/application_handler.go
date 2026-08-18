package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sai-mudike/job-application-tracker/internals/models"
	"github.com/sai-mudike/job-application-tracker/internals/repositories"
)

func CreateApplication(context *gin.Context) {
	userID := context.GetInt64("userID")

	var application models.Application

	err := context.ShouldBindJSON(&application)

	if err != nil {
		context.JSON(http.StatusBadRequest, models.NewErrorResponse("INVALID_REQUEST", "Invalid request"))
		return
	}

	application.User_id = userID

	err = repositories.CreateApplication(&application)
	if err != nil {
		HandleError(context, err)
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "application created successfully", "application": application})

}

func GetApplications(context *gin.Context) {
	userID := context.GetInt64("userID")

	applications, err := repositories.GetAllApplications(userID)
	if err != nil {
		HandleError(context, err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"applications": applications})

}

func GetApplicationByID(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	userID := context.GetInt64("userID")

	if err != nil {
		context.JSON(http.StatusBadRequest, models.NewErrorResponse("INVALID_REQUEST", "Invalid request"))
		return
	}
	application, err := repositories.GetApplicationByID(id, userID)
	if err != nil {
		HandleError(context, err)
		return
	}
	context.JSON(http.StatusOK, gin.H{"application": application})
}

func UpdateApplication(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	userID := context.GetInt64("userID")

	if err != nil {
		context.JSON(http.StatusBadRequest, models.NewErrorResponse("INVALID_REQUEST", "Invalid request"))
		return
	}
	applicationFromDB, err := repositories.GetApplicationByID(id, userID)
	if err != nil {
		HandleError(context, err)
		return
	}
	var Updatedapplication models.Application

	err = context.ShouldBindJSON(&Updatedapplication)

	if err != nil {
		context.JSON(http.StatusBadRequest, models.NewErrorResponse("INVALID_REQUEST", "Invalid request"))
		return
	}

	Updatedapplication.Id = applicationFromDB.Id
	err = repositories.UpdateApplication(Updatedapplication, id)
	if err != nil {
		HandleError(context, err)
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event updated successfully"})

}

func DeleteApplication(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	userID := context.GetInt64("userID")

	if err != nil {
		context.JSON(http.StatusBadRequest, models.NewErrorResponse("INVALID_REQUEST", "Invalid request"))
		return
	}
	_, err = repositories.GetApplicationByID(id, userID)
	if err != nil {
		HandleError(context, err)
		return
	}

	err = repositories.DeleteApplication(id)
	if err != nil {
		HandleError(context, err)
		return
	}

	context.JSON(http.StatusNoContent, gin.H{"message": "application deleted"})

}
