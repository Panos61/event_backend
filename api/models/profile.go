package models

import (
	"github.com/jinzhu/gorm"
)

// Profiles struct
type Profiles struct {
	// Owner of the profile => UserID
	UserID uint32 `gorm:"not null" json:"user_id"`
	User   User   `json:"user"`

	ID           uint64 `gorm:"primary_key;auto_increment" json:"id"`
	FirstName    string `gorm:"type:varchar(45)" json:"firstName" valid:"length(0|22)"`
	LastName     string `gorm:"type:varchar(45)" json:"lastName" valid:"length(0|22)"`
	Introduction string `gorm:"type:varchar(200)" json:"introduction" valid:"length(0|150)"`
	Age          string `gorm:"type:int" json:"age" valid:"int"`
}

// SaveProfile => Saves Profile Data into DB
func (p *Profiles) SaveProfile(db *gorm.DB) (*Profiles, error) {
	var err error
	err = db.Debug().Model(&Profiles{}).Create(&p).Error

	if err != nil {
		return &Profiles{}, err
	}

	// Assign User as the Profile owner.
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &Profiles{}, err
		}
	}

	return p, nil
}

// UpdateProfile =>..
func (p *Profiles) UpdateProfile(db *gorm.DB, uid uint32) (*Profiles, error) {
	db = db.Debug().Model(&Profiles{}).Where("id = ?", uid).Take(&Profiles{}).UpdateColumns(
		map[string]interface{}{
			"firstName":    p.FirstName,
			"lastName":     p.LastName,
			"introduction": p.Introduction,
			"age":          p.Age,
		},
	)

	if db.Error != nil {
		return &Profiles{}, db.Error
	}

	err := db.Debug().Model(&Profiles{}).Where("id = ?", uid).Take(&p).Error
	if err != nil {
		return &Profiles{}, err
	}

	return p, nil
}

// FindAllProfiles => Get all profiles stored in DB
func (p *Profiles) FindAllProfiles(db *gorm.DB) (*[]Profiles, error) {
	var err error
	profiles := []Profiles{}
	err = db.Debug().Model(&Profiles{}).Order("name DESC").Find(&profiles).Error

	if err != nil {
		return &[]Profiles{}, err
	}

	if len(profiles) > 0 {
		for i, _ := range profiles {
			err := db.Debug().Model(&User{}).Where("id = ?", profiles[i].UserID).Take(&profiles[i].User).Error
			if err != nil {
				return &[]Profiles{}, err
			}
		}
	}
	return &profiles, nil
}

//FindProfileByID => Get profile based on ID
func (p *Profiles) FindProfileByID(db *gorm.DB, pid uint64) (*Profiles, error) {
	var err error

	err = db.Debug().Model(&Events{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Profiles{}, err
	}

	if p.ID != 0 {
		err = db.Debug().Model(&Profiles{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &Profiles{}, err
		}
	}

	return p, nil
}
