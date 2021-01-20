package repository

import (
	"github.com/khalidalhabibie/depatu/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	AddUser(model.User) (model.User, error)
	GetUser(string) (model.User, error)
	UpdateUser(model.User) (model.User, error)
	DeleteUser(model.User) (model.User, error)
	GetTask(string) ([]model.Task, error)
}

type userRepository struct {
	connection *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{
		connection: DB(),
	}
}

func (db *userRepository) AddUser(user model.User) (model.User, error) {
	return user, db.connection.Create(&user).Error
}

func (db *userRepository) GetUser(username string) (user model.User, err error) {
	return user, db.connection.First(&user, "username=?", username).Error
}

func (db *userRepository) UpdateUser(user model.User) (model.User, error) {

	if err := db.connection.First(&user, user.ID).Error; err != nil {
		return user, err
	}

	return user, db.connection.Model(&user).Updates(&user).Error

}

func (db *userRepository) DeleteUser(user model.User) (model.User, error) {
	if err := db.connection.First(&user, user.Username).Error; err != nil {
		return user, err
	}
	return user, db.connection.Delete(&user).Error
}

func (db *userRepository) GetTask(username string) (tasks []model.Task, err error) {
	return tasks, db.connection.Where("assignedto = ?", username).Set("gorm:auto_preload", true).Find(&tasks).Error
}
