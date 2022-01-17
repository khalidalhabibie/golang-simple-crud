
package model
import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Username  string `json:"username" binding:"required,min=6,max=20"`
	Addressname string `json:"addressname" binding:"required,min=6,max=20"`
	Address     string `json:"address" binding:"required"`
	Provience   string `json:"provience" binding:"required"`
	City        string `json:"city" binding:"required"`
	PostalCode  int    `json:"postalcode" binding:"required"`
}

func (Address) Tablename() string {
	return "address_1"
}
