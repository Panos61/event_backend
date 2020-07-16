package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type Upvote struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	UserID    uint32    `gorm:"not null" json:"user_id"`
	PostID    uint64    `gorm:"not null" json:"post_id"`
	UpvotedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"upvoted_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// SaveUpvote => Saves upvote entity into DB
func (u *Upvote) SaveUpvote(db *gorm.DB) (*Upvote, error) {

	// Check if the auth user has liked this post before:
	err := db.Debug().Model(&Upvote{}).Where("post_id = ? AND user_id = ?", u.PostID, u.UserID).Take(&u).Error
	if err != nil {
		if err.Error() == "record not found" {
			// The user has not liked this post before, so lets save incomming like:
			err = db.Debug().Model(&Upvote{}).Create(&u).Error
			if err != nil {
				return &Upvote{}, err
			}
		}
	} else {
		// The user has liked it before, so create a custom error message
		err = errors.New("double like")
		return &Upvote{}, err
	}
	return u, nil
}

// DeleteUpvote => ..
func (u *Upvote) DeleteUpvote(db *gorm.DB) (*Upvote, error) {
	var err error
	var deletedUpvote *Upvote

	err = db.Debug().Model(Upvote{}).Where("id = ?", u.ID).Take(&u).Error
	if err != nil {
		return &Upvote{}, err
	} else {
		deletedUpvote = u
		db = db.Debug().Model(&Upvote{}).Where("id = ?", u.ID).Take(&Upvote{}).Delete(&Upvote{})

		if db.Error != nil {
			return &Upvote{}, db.Error
		}

		return deletedUpvote, nil
	}
}

func (u *Upvote) GetUpvotesInfo(db *gorm.DB, pid uint64) (*[]Upvote, error) {

	upvotes := []Upvote{}
	err := db.Debug().Model(&Upvote{}).Where("post_id = ?", pid).Find(&upvotes).Error
	if err != nil {
		return &[]Upvote{}, err
	}
	return &upvotes, err
}

// DeleteUserUpvotes => Delete user upvotes when the post gets deleted
func (u *Upvote) DeleteUserUpvotes(db *gorm.DB, uid uint32) (int64, error) {
	upvotes := []Upvote{}
	db = db.Debug().Model(&Upvote{}).Where("user_id = ?", uid).Find(&upvotes).Delete(&upvotes)
	if db.Error != nil {
		return 0, db.Error
	}

	return db.RowsAffected, nil
}

// DeleteAllPostUpvotes => Delete all upvotes when the post gets deleted
func (u *Upvote) DeleteAllPostUpvotes(db *gorm.DB, pid uint64) (int64, error) {
	upvotes := []Upvote{}
	db = db.Debug().Model(&Upvote{}).Where("post_id = ?", pid).Find(&upvotes).Delete(&upvotes)

	if db.Error != nil {
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
