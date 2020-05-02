package models

import (
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// Profile struct
type Profile struct {
	ID           uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Nickname     string    `gorm:"type:varchar(20)" json:"name" valid:"length(0|13)"`
	Introduction string    `gorm:"type:varchar(200)" json:"introduction" valid:"length(0|150)"`
	Age          string    `json:"age" valid:"int"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// SaveProfile => Inserts new query in DB
func (u *Profile) SaveProfile(db *gorm.DB) (*Profile, error) {
	var err error
	err = db.Debug().Create(&u).Error

	if err != nil {
		return &Profile{}, err
	}

	return u, nil
}

// Prepare //
func (u *Profile) Prepare() {
	u.Nickname = html.EscapeString(strings.TrimSpace(u.Nickname))
	u.Introduction = html.EscapeString(strings.TrimSpace(u.Introduction))
	//u.Age = html.EscapeString(strings.TrimSpace(u.Age))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}
