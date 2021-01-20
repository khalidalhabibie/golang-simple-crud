package model

type Task struct {
	taskname   string `json:"taskname" binding:"required,min=6,max=20" gorm:"unique"`
	assignedto string `json:"assignedto binding:"min=6,max =20"`
	status     string `json:"assignedto" binding:"required,max=1024"`
}

func (Task) TableName() string {
	return "task_1"
}
