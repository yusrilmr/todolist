package models

import (
	"gorm.io/gorm"
	"time"
)
type (
	User struct {
		gorm.Model
		FirstName string `json:"firstname"`
		LastName string `json:"lastname"`
		Email string `json:"email"`
		Password string `json:"password,omitempty"`
		HashPassword []byte `json:"hashpassword,omitempty"`
		Task []Task
	}
	Task struct {
		gorm.Model
		Name string `json:"name"`
		Description string `json:"description"`
		Due time.Time `json:"due"`
		Status string `json:"status"`
		UserID uint
		Label []*Label `gorm:"many2many:task_label;"`
	}
	Label struct {
		gorm.Model
		Description string `json:"description"`
		Task []*Task `gorm:"many2many:task_label;"`
	}
)