package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Profiles struct
type Profiles struct {
	// Owner of the profile => UserID
	UserID uint32 `gorm:"not null;primary_key;auto_increment" json:"user_id"`
	User   User   `json:"user"`

	//ID           uint64    `json:"id"` //gorm:"primary_key;auto_increment"
	FirstName    string    `gorm:"type:varchar(45)" json:"firstName" valid:"length(0|22), optional"`
	LastName     string    `gorm:"type:varchar(45)" json:"lastName" valid:"length(0|22), optional"`
	Introduction string    `gorm:"type:varchar(300)" json:"introduction" valid:"length(0|250), optional"`
	Age          string    `gorm:"type:varchar(10)" json:"age" valid:"int, optional"`
	Location     string    `gorm:"type:varchar(40)" json:"location" valid:"length(0|40), optional"`
	SharedLink   string    `gorm:"type:varchar(100)" json:"sharedLink" valid:"optional"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at" valid:"optional"`
}

// SaveProfile => Saves Profile Data into DB
func (p *Profiles) SaveProfile(db *gorm.DB) (*Profiles, error) {
	var err error

	err = db.Debug().Model(&Profiles{}).Create(&p).Error

	if err != nil {
		return &Profiles{}, err
	}

	// Assign User as the Profile owner.

	err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
	if err != nil {
		return &Profiles{}, err
	}

	return p, nil
}

// UpdateProfile =>..
func (p *Profiles) UpdateProfile(db *gorm.DB) (*Profiles, error) {
	var err error

	if p.UserID != 0 {
		err = db.Debug().Model(&Profiles{}).Where("user_id = ?", p.UserID).Updates(Profiles{FirstName: p.FirstName, LastName: p.LastName, Introduction: p.Introduction, Age: p.Age, Location: p.Location, SharedLink: p.SharedLink, UpdatedAt: time.Now()}).Error
		if err != nil {
			return &Profiles{}, err
		}
	}

	if p.UserID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &Profiles{}, err
		}
	}

	return p, nil
}

// FindAllProfiles => Get all profiles stored in DB
// func (p *Profiles) FindAllProfiles(db *gorm.DB) (*[]Profiles, error) {
// 	var err error
// 	profiles := []Profiles{}
// 	err = db.Debug().Model(&Profiles{}).Order("name DESC").Find(&profiles).Error

// 	if err != nil {
// 		return &[]Profiles{}, err
// 	}

// 	if len(profiles) > 0 {
// 		for i, _ := range profiles {
// 			err := db.Debug().Model(&User{}).Where("id = ?", profiles[i].UserID).Take(&profiles[i].User).Error
// 			if err != nil {
// 				return &[]Profiles{}, err
// 			}
// 		}
// 	}
// 	return &profiles, nil
// }

//FindProfileByID => Get profile based on ID
func (p *Profiles) FindProfileByID(db *gorm.DB, pid uint64) (*Profiles, error) {
	var err error

	err = db.Debug().Model(&Profiles{}).Where("user_id = ?", pid).Take(&p).Error
	if err != nil {
		return &Profiles{}, err
	}

	if p.UserID != 0 {
		err = db.Debug().Model(&Profiles{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &Profiles{}, err
		}
	}

	return p, nil
}
