package models

import (
	"errors"
	"event_backend/api/security"

	"github.com/jinzhu/gorm"
)

// FindUsers ..
func (u *User) FindUsers(db *gorm.DB) (*[]User, error) {

	var err error
	users := []User{}

	err = db.Debug().Model(&User{}).Find(&users).Error

	if err != nil {
		return &[]User{}, err
	}

	return &users, err

}

// FindUserByID ..
func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error

	if err != nil {
		return &User{}, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User not found")
	}

	return u, err
}

// BeforeSave .
func (u *User) BeforeSave() error {
	hashedPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}
