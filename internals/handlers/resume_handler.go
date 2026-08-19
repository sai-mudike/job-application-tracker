package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	appErrors "github.com/sai-mudike/job-application-tracker/internals/errors"
	"github.com/sai-mudike/job-application-tracker/internals/handlers/validators"
	"github.com/sai-mudike/job-application-tracker/internals/models"
	"github.com/sai-mudike/job-application-tracker/internals/repositories"
)

func CreateResume(context *gin.Context) {
	user_id := context.GetInt64("userID")
	var resume models.Resume

	err := context.ShouldBindJSON(&resume)
	if err != nil {
		HandleError(context, appErrors.ErrInvalidRequest)
		return
	}

	err = validators.ValidateCreateResume(&resume)
	if err != nil {
		fmt.Println(err)
		HandleError(context, err)
		return
	}
	resume.UserID = user_id

	err = repositories.CreateResume(resume)
	if err != nil {
		HandleError(context, err)
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "resume saved successfully", "resume": resume})
}

func GetResumes(context *gin.Context) {
	user_id := context.GetInt64("userID")
	resumes, err := repositories.GetAllResumes(user_id)
	if err != nil {
		HandleError(context, err)
		return
	}
	context.JSON(http.StatusOK, gin.H{"resumes": resumes})

}

func GetResumeByID(context *gin.Context) {
	user_id := context.GetInt64("userID")
	resume_id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		HandleError(context, appErrors.ErrResumeNotFound)
		return
	}

	resume, err := repositories.GetResumeByID(resume_id, user_id)

	if err != nil {
		HandleError(context, err)
		return
	}
	context.JSON(http.StatusOK, gin.H{"resume": resume})

}

func DeleteResume(context *gin.Context) {
	user_id := context.GetInt64("userID")
	resume_id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		HandleError(context, appErrors.ErrResumeNotFound)
		return
	}
	_, err = repositories.GetResumeByID(resume_id, user_id)

	if err != nil {
		HandleError(context, err)
		return
	}

	err = repositories.DeleteResume(resume_id, user_id)

	if err != nil {
		HandleError(context, err)
	}
	context.JSON(http.StatusOK, gin.H{"message": "Resume deleted Successfully"})

}
