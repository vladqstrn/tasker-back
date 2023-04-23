package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vladqstrn/tasker-back/internal/mw"
	"github.com/vladqstrn/tasker-back/internal/tasker/delivery/rest/handler"
	"github.com/vladqstrn/tasker-back/internal/tasker/repository"
)

func TaskRoutes(router *gin.Engine, taskRepo repository.PostgresTaskRepository) {
	taskHandler := handler.NewTaskHandler(&taskRepo)
	tasksGroup := router.Group("/tasks")

	tasksGroup.Use(mw.ProxyMiddleware())
	tasksGroup.POST("/create", taskHandler.CreateTaskHandler)
	tasksGroup.PUT("/update/:id", taskHandler.UpdateTaskHandler)
	tasksGroup.DELETE("/delete/:id", taskHandler.DeleteTaskHandler)
	tasksGroup.GET("/getalltask", taskHandler.GetAllTasksHandler)
	tasksGroup.GET("/gettask/:id", taskHandler.GetTaskByIdHandler)
	tasksGroup.GET("/getuser/:id", taskHandler.GetUserTasksHandler)
}
