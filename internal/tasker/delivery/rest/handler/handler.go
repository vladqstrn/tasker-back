package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vladqstrn/tasker-back/internal/models"
	"github.com/vladqstrn/tasker-back/internal/tasker/usecase"
)

// TaskHandler - структура обработчика задач
type TaskHandler struct {
	useCase usecase.Task
}

// NewTaskHandler возвращает новую структуру TaskHandler
func NewTaskHandler(useCase usecase.Task) *TaskHandler {
	return &TaskHandler{useCase: useCase}
}

// CreateTaskHandler - обработчик запроса для создания новой задачи
// HTTP метод POST /tasks/create
func (th *TaskHandler) CreateTaskHandler(c *gin.Context) {
	task := &models.Task{}
	if err := c.BindJSON(task); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := th.useCase.CreateTask(task); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"task": task})
}

// UpdateTaskHandler - обработчик запроса для обновления информации о существующей задаче
// HTTP метод PUT /tasks/update/:id
func (th *TaskHandler) UpdateTaskHandler(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task := &models.Task{}
	if err := c.BindJSON(task); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := th.useCase.UpdateTask(taskId, task); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": task})
}

// DeleteTaskHandler - обработчик запроса для удаления задачи
// HTTP метод DELETE /tasks/delete/:id
func (th *TaskHandler) DeleteTaskHandler(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	if err := th.useCase.DeleteTask(taskId); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetAllTasksHandler - обработчик запроса на получение списка всех задач
// HTTP метод GET /tasks/getalltask
func (th *TaskHandler) GetAllTasksHandler(c *gin.Context) {
	tasks, err := th.useCase.GetAllTasks()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// GetTaskByIdHandler - обработчик запроса на получение задачи по ее идентификатору
// HTTP метод GET /tasks/gettask/:id
func (th *TaskHandler) GetTaskByIdHandler(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, err := th.useCase.GetTaskById(taskId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": task})
}

// GetUserTasksHandler - обработчик запроса на получение списка задач пользователя по его идентификатору
// HTTP метод GET /users/getuser/:id
func (th *TaskHandler) GetUserTasksHandler(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	tasks, err := th.useCase.GetUserTasks(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user tasks"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}
