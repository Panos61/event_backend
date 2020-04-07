package models

import (
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// User struct
type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Email     string    `gorm:"size:100;not null;unique" json:"email" valid:"email"`
	Username  string    `gorm:"size:200;not null;unique" json:"username" valid:"(3|13)"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	Gender    string    `gorm:"not null;default: null" json:"gender"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

//SaveUser => Inserts new query in DB
func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error

	if err != nil {
		return &User{}, err
	}

	return u, nil
}

// Prepare //
func (u *User) Prepare() {
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}
