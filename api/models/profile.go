package models

import "time"

// Profile struct
type Profile struct {
	ID           uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Nickname     string    `gorm:"type:varchar(20)" json:"name" valid:"length(0|13)"`
	Intruduction string    `gorm:"type:varchar(200)" json:"introduction" valid:"length(0|150)"`
	Age          uint      `json:"age"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
