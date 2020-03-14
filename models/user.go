package models

import (
	"github.com/jinzhu/gorm"
)

//User struct declaration
type User struct {
	gorm.Model

	Username string `gorm:"type:varchar(18);unique_index;not null;default:null" json:"username" valid:"length(3|13)"`
	Email    string `gorm:"type:varchar(100);unique_index;not null;default:null " json:"email" valid:"email"`
	Gender   string `json:"gender" gorm:"not null;default: null"`
	Password string `json:"password" gorm:"not null;default: null" valid:"length(8|100)"`
}
