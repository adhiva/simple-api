package controllers

import (
	"net/http"
	"sample-api/db"

	"github.com/gin-gonic/gin"

	"sample-api/models"
)

func GetTodos(c *gin.Context) {

	var tasks []models.Todo
	db := db.GetDB()
	db.Find(&tasks)
	c.JSON(200, tasks)
}

func CreateTodo(c *gin.Context) {
	var task models.Todo
	var db = db.GetDB()

	if err := c.BindJSON(&task); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	db.Create(&task)
	c.JSON(http.StatusOK, &task)
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var task models.Todo

	db := db.GetDB()
	if err := db.Where("id = ?", id).First(&task).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.BindJSON(&task)
	db.Save(&task)
	c.JSON(http.StatusOK, &task)
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	var task models.Todo
	db := db.GetDB()

	if err := db.Where("id = ?", id).First(&task).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	db.Delete(&task)
}
