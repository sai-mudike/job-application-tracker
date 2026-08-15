package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sai-mudike/job-application-tracker/internals/models"
	"github.com/sai-mudike/job-application-tracker/internals/repository"
)

func CreateCompany(context *gin.Context) {
	var company models.Company

	err := context.ShouldBindJSON(&company)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the given company data"})
		return
	}

	company.User_id = 1

	err = repository.SaveCompany(&company)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Company Created", "company": company})

}

func GetCompanies(context *gin.Context) {
	companies, err := repository.GetAllCompanies()

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not get the companies", "error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"companies": companies})

}

func CompanyByID(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the id provided"})
		return
	}

	company, err := repository.GetCompanyByID(id)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the data"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"company": company})

}
