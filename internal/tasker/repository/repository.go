package repository

import (
	"github.com/go-pg/pg"
	"github.com/vladqstrn/tasker-back/internal/models"
)

type PostgresTaskRepository struct {
	db *pg.DB
}

func NewPostgresTaskRepository(db *pg.DB) *PostgresTaskRepository {
	return &PostgresTaskRepository{db: db}
}

// CreateTask - добавляет новую задачу в базу данных
func (db *PostgresTaskRepository) CreateTask(task *models.Task) error {
	_, err := db.db.Model(task).
		Insert()
	if err != nil {
		return err
	}

	return nil
}

// UpdateTask - обновляет информацию о задаче
func (db *PostgresTaskRepository) UpdateTask(taskId int, task *models.Task) error {
	_, err := db.db.Model(task).
		Column("title", "text", "description", "executor", "status", "created_at").
		WherePK().
		Update()
	if err != nil {
		return err
	}

	return nil
}

// DeleteTask - удаляет задачу из БД по ее идентификатору
func (db *PostgresTaskRepository) DeleteTask(taskId int) error {
	_, err := db.db.Model(&models.Task{}).
		Where("id = ?", taskId).
		Delete()
	if err != nil {
		return err
	}

	return nil
}

// GetAllTasks - получает список всех задач из БД
func (db *PostgresTaskRepository) GetAllTasks() ([]*models.Task, error) {
	var tasks []*models.Task
	err := db.db.Model(&tasks).
		Select()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// GetTaskById - получает запись о задаче по ее ID
func (db *PostgresTaskRepository) GetTaskById(taskId int) (*models.Task, error) {
	task := &models.Task{Id: taskId}
	err := db.db.Model(task).
		WherePK().
		Select()
	if err != nil {
		return nil, err
	}

	return task, nil
}

// GetUserTasks - получает список задач пользователя с данным ID
func (db *PostgresTaskRepository) GetUserTasks(userId int) ([]string, error) {
	var descriptions []string
	_, err := db.db.Query(&descriptions, `SELECT description FROM tasks WHERE id = ?`, userId)
	if err != nil {
		return nil, err
	}

	return descriptions, nil
}
