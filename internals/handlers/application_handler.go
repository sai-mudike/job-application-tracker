package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sai-mudike/job-application-tracker/internals/handlers/validators"
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

	err = validators.ValidateCreateApplication(application)
	if err != nil {
		HandleError(context, err)
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

	query := models.ApplicationQuery{
		Page:           1,
		Limit:          10,
		Status:         context.DefaultQuery("status", ""),
		EmploymentType: context.DefaultQuery("employment_type", ""),
		SortBy:         context.DefaultQuery("sort", "created_at"),
		OrderBy:        context.DefaultQuery("order", "desc"),
	}

	err := validators.ApplQueryVali(context, &query, HandleError)
	if err != nil {
		fmt.Println(err)
		HandleError(context, err)
		return
	}

	applications, err := repositories.GetAllApplications(userID, query)
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
	ApplicationID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	userID := context.GetInt64("userID")

	if err != nil {
		context.JSON(http.StatusBadRequest, models.NewErrorResponse("INVALID_REQUEST", "Invalid request"))
		return
	}
	_, err = repositories.GetApplicationByID(ApplicationID, userID)
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
	err = validators.ValidateCreateApplication(Updatedapplication)
	if err != nil {
		HandleError(context, err)
		return
	}

	Updatedapplication.Id = ApplicationID
	err = repositories.UpdateApplication(Updatedapplication, ApplicationID)
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
