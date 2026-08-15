package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sai-mudike/job-application-tracker/internals/models"
	"github.com/sai-mudike/job-application-tracker/internals/repositories"
)

func CreateApplication(context *gin.Context) {

	var application models.Application

	err := context.ShouldBindJSON(&application)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the given data"})
		return
	}

	application.User_id = 1

	err = repositories.CreateApplication(&application)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create the application"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "application created successfully", "application": application})

}

func GetApplication(context *gin.Context) {

	applications, err := repositories.GetAllApplications()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the data"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"applications": applications})

}

func GetApplicationByID(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the given id"})
		return
	}
	application, err := repositories.GetApplicationByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the application"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"application": application})
}

func UpdateApplication(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the given data"})
		return
	}
	applicationFromDB, err := repositories.GetApplicationByID(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not find the application with given id"})
		return
	}
	var Updatedapplication models.Application

	err = context.ShouldBindJSON(&Updatedapplication)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the given data"})
		return
	}

	Updatedapplication.Id = applicationFromDB.Id
	err = repositories.UpdateApplication(Updatedapplication, id)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update the application"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event updated successfully"})

}

func DeleteApplication(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the given data"})
		return
	}
	_, err = repositories.GetApplicationByID(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not find the application with given id"})
		return
	}

	err = repositories.DeleteApplication(id)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete the application"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "application deleted"})

}
