package usecase

import (
	"github.com/vladqstrn/tasker-back/internal/models"
	"github.com/vladqstrn/tasker-back/internal/tasker/repository"
)

type Task interface {
	CreateTask(task *models.Task) error             //создает новую задачу.
	UpdateTask(taskId int, task *models.Task) error //обновляет данные по задаче с заданным идентификатором.
	DeleteTask(taskId int) error                    //удаляет запись о задаче с переданным идентификатором.
	GetAllTasks() ([]*models.Task, error)           //получает все сохраненные задачи.
	GetTaskById(taskId int) (*models.Task, error)   //получает задачу с конкретным идентификатором.
	GetUserTasks(userId int) ([]string, error)      //получает все задачи пользователя с данным идентификатором.
}

type TaskUc struct {
	task       Task
	repository repository.PostgresTaskRepository
}

func NewTaskUc(task Task) *TaskUc {
	return &TaskUc{task: task}
}

func (uc *TaskUc) CreateTask(task *models.Task) error {
	err := uc.repository.CreateTask(task)
	if err != nil {
		return err
	}
	return nil
}

func (uc *TaskUc) UpdateTask(taskId int, task *models.Task) error {
	task.Id = taskId
	err := uc.repository.UpdateTask(taskId, task)
	if err != nil {
		return err
	}
	return nil
}

func (uc *TaskUc) DeleteTask(taskId int) error {
	err := uc.repository.DeleteTask(taskId)
	if err != nil {
		return err
	}
	return nil
}

func (uc *TaskUc) GetAllTasks() ([]*models.Task, error) {
	task, err := uc.repository.GetAllTasks()
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (uc *TaskUc) GetTaskById(taskId int) (*models.Task, error) {
	task, err := uc.repository.GetTaskById(taskId)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (uc *TaskUc) GetUserTasks(userId int) ([]string, error) {
	tasks, err := uc.repository.GetUserTasks(userId)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
