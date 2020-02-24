package models

import (
	"github.com/jinzhu/gorm"
)

//User struct declaration
type User struct {
	gorm.Model

	Username string `json:"username"`
	Email    string `gorm:"type:varchar(100);unique_index" json:"email"`
	Gender   string `json:"gender"`
	Password string `json:"password"`
}
