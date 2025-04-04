package main

import (
	"github.com/DiasSatymbek05/go_lang/config"
	"github.com/DiasSatymbek05/go_lang/models"
	"github.com/DiasSatymbek05/go_lang/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	config.ConnectDatabase()

	// Migrate models
	config.DB.AutoMigrate(&models.Todo{})

	// Register routes
	routes.TodoRoutes(r)

	// Start server
	r.Run(":8080")
}
