package repository

import (
	"golang-simple-crud/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	AddUser(model.User) (model.User, error)
	GetUser(string) (model.User, error)
	UpdateUser(model.User) (model.User, error)
	UpdatePassword(model.User) (model.User, error)
	DeleteUser(model.User) (model.User, error)
	GetAllUser() ([]model.User, error)
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

func (db *userRepository) CekUser(user model.User) (status bool) {
	if err := db.connection.First(&user, user.Username).Error; err != nil {
		return false
	}
	return true
}

func (db *userRepository) UpdateUser(user model.User) (model.User, error) {
	dataUpdate := user

	if err := db.connection.First(&user, user.ID).Error; err != nil {
		return user, err
	}
	dataUpdate.Image = user.Image
	dataUpdate.Password = user.Password
	return user, db.connection.Model(&user).Updates(dataUpdate).Error

}

func (db *userRepository) UpdatePassword(user model.User) (model.User, error) {
	dataUpdate := user

	if err := db.connection.First(&user, user.ID).Error; err != nil {
		return user, err
	}
	dataUpdate.Image = user.Image
	dataUpdate.Email = user.Email
	dataUpdate.Bio = user.Bio

	return user, db.connection.Model(&user).Updates(dataUpdate).Error
}

func (db *userRepository) UpdatePhotoUser(user model.User) (model.User, error) {
	dataUpdate := user

	if err := db.connection.First(&user, user.ID).Error; err != nil {
		return user, err
	}

	return user, db.connection.Model(&user).Updates(dataUpdate).Error

}

func (db *userRepository) DeleteUser(user model.User) (model.User, error) {
	if err := db.connection.First(&user, user.Username).Error; err != nil {
		return user, err
	}
	return user, db.connection.Delete(&user).Error
}

func (db *userRepository) GetAllUser() (users []model.User, err error) {
	return users, db.connection.Find(&users).Error
}
