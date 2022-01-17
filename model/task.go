package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Taskname   string `json:"taskname" binding:"required,min=6,max=20" gorm:"unique"`
	Assignedto string `json:"assignedto" binding:"required,min=6,max=20"`
	Status     string `json:"status" binding:"required,max=1024"`
}

func (Task) TableName() string {
	return "task_1"
}
