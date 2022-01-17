package repository

import (
	"golang-simple-crud/model"

	"gorm.io/gorm"
)

type TaskRepository interface {
	AddTask(model.Task) (model.Task, error)
	GetTaskByUser(string) ([]model.Task, error)
	GetAllTask() ([]model.Task, error)
	//UpdateTask(model.Task)(model.Task, error)
	DeleteTask(model.Task) (model.Task, error)
}

type taskRepository struct {
	connection *gorm.DB
}

func NewTaskRepository() TaskRepository {
	return &taskRepository{
		connection: DB(),
	}
}

func (db *taskRepository) AddTask(task model.Task) (model.Task, error) {
	return task, db.connection.Create(&task).Error
}

func (db *taskRepository) GetTaskByUser(username string) (tasks []model.Task, err error) {
	return tasks, db.connection.Where("assignedto=?", username).Set("gorm:auto_preload", true).Find(&tasks).Error
}

func (db *taskRepository) GetAllTask() (tasks []model.Task, err error) {
	return tasks, db.connection.Find(&tasks).Error
}

func (db *taskRepository) DeleteTask(task model.Task) (model.Task, error) {
	if err := db.connection.First(&task, task.ID).Error; err != nil {
		return task, err
	}
	return task, db.connection.Delete(&task).Error
}
