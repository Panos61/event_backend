package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Events => Primary struct
type Events struct {
	//CreatorID
	CreatorID uint32 `gorm:"not null" json:"creator_id"`
	Creator   User   `json:"creator"`

	// ** *
	ID          uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Category    string `gorm:"size:25;not null" json:"category"`
	Title       string `gorm:"size:100;not null" json:"title"`
	Description string `gorm:"size:1000;not null" json:"description"`

	// DateType string `gorm:"size:15;not null" json:"dateType"`
	// Date     string `gorm:"not null" json:"date"`
	// Time     string `gorm:"not null" json:"singleTime"`

	Comments string `gorm:"size:300" json:"comments"`
	//AgeRestriction string `json:"ageRestricted"`

	// Payment string `gorm:"not null" json:"payment"`
	// Price   string `json:"price"`
	//Tickets string `gorm:"size:200" json:"tickets"`

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// SaveEvent => Saves Event into DB
func (e *Events) SaveEvent(db *gorm.DB) (*Events, error) {
	var err error
	err = db.Debug().Model(&Events{}).Create(&e).Error
	//err = db.Debug().Create(&e).Error

	if err != nil {
		return &Events{}, err
	}

	// Assign User as the event Creator.
	if e.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", e.CreatorID).Take(&e.Creator).Error
		if err != nil {
			return &Events{}, err
		}
	}

	return e, nil
}

// FindAllEvents => Get all events stored in the DB
func (e *Events) FindAllEvents(db *gorm.DB) (*[]Events, error) {
	var err error
	events := []Events{}
	err = db.Debug().Model(&Events{}).Limit(100).Order("created_at desc").Find(&events).Error

	if err != nil {
		return &[]Events{}, err
	}

	if len(events) > 0 {
		for i, _ := range events {
			err := db.Debug().Model(&User{}).Where("id = ?", events[i].CreatorID).Take(&events[i].Creator).Error
			if err != nil {
				return &[]Events{}, err
			}
		}
	}
	return &events, nil
}

// FindUserEvents => Get all events created by specific user
func (e *Events) FindUserEvents(db *gorm.DB, uid uint32) (*[]Events, error) {
	var err error
	events := []Events{}

	err = db.Debug().Model(&Events{}).Where("creator_id = ?", uid).Limit(100).Order("created_at desc").Find(&events).Error
	if err != nil {
		return &[]Events{}, err
	}

	if len(events) > 0 {
		for i, _ := range events {
			err := db.Debug().Model(&Events{}).Where("id = ?", events[i].CreatorID).Take(&events[i].Creator).Error
			if err != nil {
				return &[]Events{}, err
			}
		}
	}
	return &events, nil
}

// FindEventByID => Get event based on ID
func (e *Events) FindEventByID(db *gorm.DB, pid uint64) (*Events, error) {
	var err error

	err = db.Debug().Model(&Events{}).Where("id = ?", pid).Take(&e).Error
	if err != nil {
		return &Events{}, err
	}

	if e.ID != 0 {
		err = db.Debug().Model(&Events{}).Where("id = ?", e.CreatorID).Take(&e.Creator).Error
		if err != nil {
			return &Events{}, err
		}
	}

	return e, nil
}
