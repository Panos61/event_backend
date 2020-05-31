package models

import (
	"errors"
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
	User         User      `json:"user"`
	UserID       uint32    `gorm:"not null" json:"user_id"`
}

// SaveProfile => Inserts new query in DB
func (u *Profile) SaveProfile(db *gorm.DB) (*Profile, error) {
	var err error
	err = db.Debug().Create(&u).Error

	if err != nil {
		return &Profile{}, err
	}

	// Asign User as the owner of the profile
	if u.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", u.UserID).Take(&u.User).Error
		if err != nil {
			return &Profile{}, err
		}
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

// FindProfileByID => ..
func (u *Profile) FindProfileByID(db *gorm.DB, proID uint32) (*Profile, error) {
	var err error
	err = db.Debug().Model(Profile{}).Where("id = ?", proID).Take(&u).Error

	if err != nil {
		return &Profile{}, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return &Profile{}, errors.New("Profile Not Found")
	}

	return u, nil
}
