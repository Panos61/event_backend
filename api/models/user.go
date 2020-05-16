package models

import (
	"html"
	"log"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// User struct
type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Username  string    `gorm:"size:100;not null;unique" json:"username" `
	Password  string    `gorm:"size:200;not null;" json:"password"`
	Gender    string    `gorm:"not null;" json:"gender"`
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

// DeleteAUser => ..
func (u *User) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//UpdateUserPassword => Replaces old user password with new one
func (u *User) UpdateUserPassword(db *gorm.DB, uid uint32) (*User, error) {

	if u.Password != "" {
		// Hash the new password
		err := u.BeforeSave()
		if err != nil {
			log.Fatal(err)
		}

		// Update columns
		db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
			map[string]interface{}{
				"password":   u.Password,
				"updated_at": time.Now(),
			},
		)
	}

	// If error, return error
	if db.Error != nil {
		return &User{}, db.Error
	}

	err := db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}

	return u, nil
}

// UpdateUserEmail => Replaces old email with a new one
func (u *User) UpdateUserEmail(db *gorm.DB, uid uint32) (*User, error) {
	if u.Password != "" {
		// Update Columns
		db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
			map[string]interface{}{
				"email":      u.Email,
				"updated_at": time.Now(),
			},
		)
	}

	// If error, return error
	if db.Error != nil {
		return &User{}, db.Error
	}

	err := db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}

	return u, nil
}
