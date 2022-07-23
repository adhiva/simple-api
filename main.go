package main

import (
	"log"
	"sample-api/db"
	"time"

	TodosController "sample-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	port := ":5000"
	log.Println("Starting server..")

	db.Init()

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		tasks := v1.Group("/todos")
		{
			tasks.GET("/", TodosController.GetTodos)
			tasks.POST("/", TodosController.CreateTodo)
			tasks.PUT("/:id", TodosController.UpdateTodo)
			tasks.DELETE("/:id", TodosController.DeleteTodo)
		}

		general := v1.Group("/general")
		{
			type HealthCheck struct {
				Status    string `json:"status"`
				TimeStamp string `json:"timestamp"`
			}

			general.GET("/health-check", func(c *gin.Context) {
				c.JSON(200, HealthCheck{
					Status:    "Healthy",
					TimeStamp: time.Now().Format(time.RFC3339),
				})
			})
		}
	}

	r.Run(port)
}
