package repository

import (
	"golang-simple-crud/model"

	"gorm.io/gorm"
)

type AddressRepository interface {
	AddAddress(model.Address) (model.Address, error)
	GetAddressByUsername(string) ([]model.Address, error)
	UpdateAddress(model.Address) (model.Address, error)
	DeleteAddress(model.Address) (model.Address, error)
}

type addressRepository struct {
	connection *gorm.DB
}

func NewAddressRepository() AddressRepository {
	return &addressRepository{
		connection: DB(),
	}
}

func (db *addressRepository) AddAddress(address model.Address) (model.Address, error) {
	return address, db.connection.Create(&address).Error
}

func (db *addressRepository) GetAddressByUsername(username string) (address []model.Address, err error) {
	return address, db.connection.Where("username=?", username).Set("gorm:auto_preload", true).Find(&address).Error
}

func (db *addressRepository) UpdateAddress(address model.Address) (model.Address, error) {
	dataUpdate := address

	if err := db.connection.First(&address, address.ID).Error; err != nil {
		return address, err
	}

	if address.Username != dataUpdate.Username {
		address.Username = ""
		address.Addressname = ""
		address.Provience = ""
		address.Address = ""
		address.City = ""
		address.PostalCode = 0
		return address, nil
	}
	return address, db.connection.Model(&address).Updates(dataUpdate).Error

}

func (db *addressRepository) DeleteAddress(address model.Address) (model.Address, error) {
	dataDeleted := address
	if err := db.connection.First(&address, address.ID).Error; err != nil {
		return address, err
	}

	if address.Username != dataDeleted.Username {
		address.Username = ""
		address.Addressname = ""
		address.Provience = ""
		address.Address = ""
		address.City = ""
		address.PostalCode = 0
		return address, nil
	}
	return address, db.connection.Delete(&address).Error
}
