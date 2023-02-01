package main

import (
  "github.com/gin-gonic/gin"
  "github.com/rest-api-go/models"
  "github.com/rest-api-go/controllers"
)

func main() {
  r := gin.Default()
  models.ConnectDatabase()

  r.GET("/api/tasks", controllers.FindTasks)
  r.POST("/api/tasks", controllers.CreateTask)
  r.GET("/api/tasks/one", controllers.FindTask) 
  r.PUT("/api/tasks/update", controllers.UpdateTask) 
  r.Run()
}