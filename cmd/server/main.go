package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sai-mudike/job-application-tracker/internals/database"
	"github.com/sai-mudike/job-application-tracker/internals/handlers"
)

func main() {

	server := gin.Default()
	database.Init()
	server.GET("/applications", handlers.GetApplication)
	server.GET("/applications/:id", handlers.GetApplicationByID)
	server.POST("/applications", handlers.CreateApplication)
	server.PUT("/applications/:id", handlers.UpdateApplication)
	server.DELETE("applications/:id", handlers.DeleteApplication)

	server.POST("/auth/register", handlers.RegisterUser)
	server.POST("/auth/login", handlers.UserLogin)

	server.Run(":8080")
}
