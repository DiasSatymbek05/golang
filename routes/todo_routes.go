package routes

import (
	"github.com/DiasSatymbek05/go_lang/controllers"

	"github.com/gin-gonic/gin"
)

func TodoRoutes(router *gin.Engine) {

	router.GET("/todos", controllers.GetTodos)
	router.GET("/todos/:id", controllers.GetTodoByID)
	router.POST("/todos", controllers.CreateTodo)
	router.PUT("/todos/:id", controllers.UpdateTodo)
	router.DELETE("/todos/:id", controllers.DeleteTodo)
}
