package repository

import (
	"github.com/khalidalhabibie/depatu/model"
	"gorm.io/gorm"
)

type TaskRepository interface {
	GetTaskByUser(string) ([]model.Task, error)
	GetAllTask() ([]model.Task, error)
	//UpdateTask(model.Task)(model.Task, error)
	//DeleteTask(string) (model.Task, error)
}

type taskRepository struct {
	connection *gorm.DB
}

func NewTaskRepository() TaskRepository {
	return &taskRepository{
		connection: DB(),
	}
}

func (db *taskRepository) GetTaskByUser(username string) (tasks []model.Task, err error) {
	return tasks, db.connection.Where("assignedto = ?", username).Set("gorm:auto_preload", true).Find(&tasks).Error
}

func (db *taskRepository) GetAllTask() (tasks []model.Task, err error) {
	return tasks, db.connection.Find(&tasks).Error
}

/*
func (db *taskRepository) UpdateTask(task model.Task) (model.Task, error) {
	if err := db.connection.First(&task, task.taskname).Error; err != nil {
		return task, err
	}
	return task, db.connection.Model(&task).Updates(&task).Error
}


func (db *taskRepository) DeleteProduct(task model.Task) (model.Task, error) {
	if err := db.connection.First(&task, task.taskname).Error; err != nil {
		return task, err
	}
	return task, db.connection.Delete(&task).Error
}

*/
