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
	Description string `gorm:"not null" json:"description"`
	Avatar      string `gorm:"not null" json:"avatar"`

	City    string `gorm:"size: 40;not null" json:"city"`
	Address string `gorm:"size: 50;not null" json:"address"`
	Place   string `gorm:"size: 45;not null" json:"place"`

	DateType string `gorm:"size: 30;not null" json:"dateType"`
	Date     string `gorm:"size: 100;not null" json:"date"`
	Time     string `gorm:"size: 100;not null" json:"singleTime"`

	Comments   string `gorm:"size:200" json:"comments"`
	URLYoutube string `gorm:"size:200" json:"urlYoutube"`
	//AgeRestriction string `json:"ageRestricted"`

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

// DeleteUserEvents => Delete all user's events
func (e *Events) DeleteUserEvents(db *gorm.DB, uid uint32) (int64, error) {
	events := []Events{}
	db = db.Debug().Model(&Events{}).Where("creator_id = ?", uid).Find(&events).Delete(&events)

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// GetMusic => Fetches all music events
func (e *Events) GetMusic(db *gorm.DB) (*[]Events, error) {
	var err error
	events := []Events{}
	var category string = "Μουσική"

	err = db.Debug().Model(&Events{}).Where("category = ?", category).Order("created_at desc").Find(&events).Error
	if err != nil {
		return &[]Events{}, err
	}

	return &events, nil
}

//GetSports => Fetches all sport evevts
func (e *Events) GetSports(db *gorm.DB) (*[]Events, error) {
	var err error
	events := []Events{}
	var category string = "Αθλητισμός"

	err = db.Debug().Model(&Events{}).Where("category = ?", category).Order("created_at desc").Find(&events).Error
	if err != nil {
		return &[]Events{}, err
	}

	return &events, nil
}

//GetEnter => Fetches all entertainment evevts
func (e *Events) GetEnter(db *gorm.DB) (*[]Events, error) {
	var err error
	events := []Events{}
	var category string = "Διασκέδαση"

	err = db.Debug().Model(&Events{}).Where("category = ?", category).Order("created_at desc").Find(&events).Error
	if err != nil {
		return &[]Events{}, err
	}

	return &events, nil
}

//GetCinema => Fetches all cinema evevts
func (e *Events) GetCinema(db *gorm.DB) (*[]Events, error) {
	var err error
	events := []Events{}
	var category string = "Σινεμά"

	err = db.Debug().Model(&Events{}).Where("category = ?", category).Order("created_at desc").Find(&events).Error
	if err != nil {
		return &[]Events{}, err
	}

	return &events, nil
}

//GetArts => Fetches all art evevts
func (e *Events) GetArts(db *gorm.DB) (*[]Events, error) {
	var err error
	events := []Events{}
	var category string = "Τέχνες"

	err = db.Debug().Model(&Events{}).Where("category = ?", category).Order("created_at desc").Find(&events).Error
	if err != nil {
		return &[]Events{}, err
	}

	return &events, nil
}

//GetDance => Fetches all dance evevts
func (e *Events) GetDance(db *gorm.DB) (*[]Events, error) {
	var err error
	events := []Events{}
	var category string = "Χορός"

	err = db.Debug().Model(&Events{}).Where("category = ?", category).Order("created_at desc").Find(&events).Error
	if err != nil {
		return &[]Events{}, err
	}

	return &events, nil
}

//GetKids => Fetches all kids' evevts
func (e *Events) GetKids(db *gorm.DB) (*[]Events, error) {
	var err error
	events := []Events{}
	var category string = "Παιδικά"

	err = db.Debug().Model(&Events{}).Where("category = ?", category).Order("created_at desc").Find(&events).Error
	if err != nil {
		return &[]Events{}, err
	}

	return &events, nil
}

//GetSocial => Fetches all social events
func (e *Events) GetSocial(db *gorm.DB) (*[]Events, error) {
	var err error
	events := []Events{}
	var category string = "Social_Events"

	err = db.Debug().Model(&Events{}).Where("category = ?", category).Order("created_at desc").Find(&events).Error
	if err != nil {
		return &[]Events{}, err
	}

	return &events, nil
}
