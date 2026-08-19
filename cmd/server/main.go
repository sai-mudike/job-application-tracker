package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sai-mudike/job-application-tracker/internals/database"
	"github.com/sai-mudike/job-application-tracker/internals/handlers"
	middelware "github.com/sai-mudike/job-application-tracker/internals/middleware"
)

func main() {

	server := gin.Default()
	database.Init()

	authenticater := server.Group("/")
	authenticater.Use(middelware.Authentication)

	// Applications route
	authenticater.GET("/applications", handlers.GetApplications)
	authenticater.GET("/applications/:id", handlers.GetApplicationByID)
	authenticater.POST("/applications", handlers.CreateApplication)
	authenticater.PUT("/applications/:id", handlers.UpdateApplication)
	authenticater.DELETE("applications/:id", handlers.DeleteApplication)

	//  Resumes route
	authenticater.GET("/resume", handlers.GetResumes)
	authenticater.GET("/resume/:id", handlers.GetResumeByID)
	authenticater.POST("/resume", handlers.CreateResume)
	authenticater.DELETE("/resume/:id", handlers.DeleteResume)

	// users route
	server.POST("/auth/register", handlers.RegisterUser)
	server.POST("/auth/login", handlers.UserLogin)

	server.Run(":8080")
}
