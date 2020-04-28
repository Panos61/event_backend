package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Events => Primary struct
type Events struct {
	ID          uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Category    string `gorm:"size:25;not null" json:"category"`
	Title       string `gorm:"size:100;not null" json:"title"`
	Description string `gorm:"size:1000;not null" json:"description"`

	DateType string `gorm:"size:15;not null" json:"dateType"`
	Date     uint32 `gorm:"not null" json:"date"`
	Time     uint32 `gorm:"not null" json:"singleTime"`

	Comments       string `gorm:"size:300" json:"comments"`
	AgeRestriction bool   `json:"ageRestricted"`

	Payment bool   `gorm:"not null" json:"payment"`
	Price   uint   `json:"price"`
	Tickets string `gorm:"size:200" json:"tickets"`

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// SaveEvent => Saves Event into DB
func (e *Events) SaveEvent(db *gorm.DB) (*Events, error) {
	var err error
	err = db.Debug().Model(&Events{}).Create(&e).Error

	if err != nil {
		return &Events{}, err
	}

	return e, nil
}
