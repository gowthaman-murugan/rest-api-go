package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rest-api-go/models"
)

type CreateTaskInput struct {
	TaskName   string `json:"task_name" binding:"required"`
	TaskDetail string `json:"task_detail" binding:"required"`
}

type UpdateTaskInput struct {
	TaskName   string `json:"task_name" binding:"required"`
	TaskDetail string `json:"task_detail" binding:"required"`
}

func FindTasks(c *gin.Context) {
	var tasks []models.Task
	models.DB.Find(&tasks)
	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func CreateTask(c *gin.Context) {
	// Validate input
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create task
	task := models.Task{TaskName: input.TaskName, TaskDetail: input.TaskDetail, Date: "2020-03-10 00:00:00"}
	models.DB.Create(&task)

	c.JSON(http.StatusOK, gin.H{"data": task})

}

func FindTask(c *gin.Context) { // Get model if exist
	var task models.Task
	id := c.Request.URL.Query().Get("id")
	if err := models.DB.Where("id = ?", id).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})

}

func UpdateTask(c *gin.Context) {
	// Get model if exist
	var task models.Task
	id := c.Request.URL.Query().Get("id")
	if err := models.DB.Where("id = ?", id).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Validate input
	var input UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	update_blog := models.Task{TaskName: input.TaskName, TaskDetail: input.TaskDetail}
	models.DB.Model(&task).Updates(&update_blog)
	c.JSON(http.StatusOK, gin.H{"data": task})
}
