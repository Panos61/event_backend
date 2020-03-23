package models

import (
	"github.com/jinzhu/gorm"
)

// ProfileSecData => Profile secondary info struct
type ProfileSecData struct {
	gorm.Model

	Nickname     string `json:"name" gorm:"type:varchar(18);unique_index;default:'@-'" valid:"length(0|13)"`
	Intruduction string `json:"introduction" gorm:"type:varchar(200);default:'Δεν υπάρχει περιγραφή.'" valid:"length(0|150)"`
	Age          int    `json:"age"`
}
