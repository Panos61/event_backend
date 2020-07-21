package models

import (
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Event_Comment struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	UserID    uint32    `gorm:"not null" json:"user_id"`
	EventID   uint64    `gorm:"not null" json:"event_id"`
	Content   string    `gorm:"text;not null" json:"content"`
	User      User      `json:"user"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (c *Event_Comment) Prepare() {
	c.ID = 0
	c.User = User{}
	c.Content = html.EscapeString(strings.TrimSpace(c.Content))
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
}

// func (e *Event_Comment) Validate(action string) map[string]string {
// 	var errorMessages = make(map[string]string)
// 	var err error

// }

// SaveEventComment = > Saves Event_Comments into the DB
func (c *Event_Comment) SaveEventComment(db *gorm.DB) (*Event_Comment, error) {
	err := db.Debug().Create(&c).Error
	if err != nil {
		return &Event_Comment{}, err
	}

	// Assign User as the owner of the comment
	if c.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", c.UserID).Take(&c.User).Error
		if err != nil {
			return &Event_Comment{}, nil
		}
	}
	return c, nil
}

func (c *Event_Comment) GetEventComments(db *gorm.DB, event_id uint64) (*[]Event_Comment, error) {
	eventComments := []Event_Comment{}

	err := db.Debug().Model(&Event_Comment{}).Where("event_id = ?", event_id).Order("created_at desc").Find(&eventComments).Error
	if err != nil {
		return &[]Event_Comment{}, err
	}

	if len(eventComments) > 0 {
		for i, _ := range eventComments {
			err := db.Debug().Model(&User{}).Where("id = ?", eventComments[i].UserID).Take(&eventComments[i].User).Error
			if err != nil {
				return &[]Event_Comment{}, err
			}
		}
	}

	return &eventComments, err
}

// DeleteUserEventComments => When a user deletes his account, his comments are getting deleted too
func (c *Event_Comment) DeleteUserEventComments(db *gorm.DB, uid uint32) (int64, error) {
	eventComments := []Event_Comment{}

	db = db.Debug().Model(&Event_Comment{}).Where("user_id = ?", uid).Find(&eventComments).Delete(&eventComments)
	if db.Error != nil {
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
