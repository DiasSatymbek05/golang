package controllers

import (
	"net/http"

	"github.com/DiasSatymbek05/go_lang/config"
	"github.com/DiasSatymbek05/go_lang/models"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	config.DB.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func GetTodoByID(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	if err := config.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Задача не найдена"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&todo)
	c.JSON(http.StatusCreated, todo)
}

func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	if err := config.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	if err := config.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	config.DB.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
