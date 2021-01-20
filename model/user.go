package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string    `json:"username" binding:"alphanum,min=4,max=255" gorm:"unique"`
	Email    string    `json:"email" binding:"required" gorm:"unique"`
	Password string    `json:"password" binding:"required,min=6,max=20"`
	Image    string    `form:"image" json:"image" binding:"omitempty,url"`
	Bio      string    `form:"bio" json:"bio" binding:"max=1024"`
	Addresss []Address `json:"address" gorm:"foreignKey:UserRefer" `
}

type Address struct {
	gorm.Model
	UserRefer   string `json:"userrefer"`
	Addressname string `json:useraddress`
	Address     string `json:"address"`
	Provience   string `json:"provience"`
	City        string `json:"city"`
	PostalCode  int    `json:"postalcode"`
}

func (User) TableName() string {
	return "user_1"
}

func (Address) Tablename() string {
	return "address_1"
}
