package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sai-mudike/job-application-tracker/internals/handlers"
	"github.com/sai-mudike/job-application-tracker/internals/migrations"
)

func main() {

	server := gin.Default()
	migrations.InitDB()

	server.GET("/companies", handlers.GetCompanies)
	server.GET("/companies/:id", handlers.CompanyByID)
	server.POST("/companies", handlers.CreateCompany)

	server.Run(":8080")
}
