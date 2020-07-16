package models

import (
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// Post => Post main struct for Feed
type Post struct {
	// CreatorID
	AuthorID uint32 `gorm:"not null" json:"author_id"`
	Author   User   `json:"author"`

	// * **
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Content   string    `gorm:"text;not null" json:"content"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Post) Prepare() {
	p.Content = html.EscapeString(strings.TrimSpace(p.Content))
	p.Author = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

// func (p *Post) Validate() map[string]string {
// 	var err error
// 	var errorMessages = make(map[string]string)

// 	if p.Content == "" {
// 		err = errors.New("Required Content")
// 		errorMessages["Required_content"] = err.Error()
// 	}

// 	if p.AuthorID < 1 {
// 		err = errors.New("Required Author")
// 		errorMessages["Required_author"] = err.Error()
// 	}
// 	return errorMessages
// }

func (p *Post) SavePost(db *gorm.DB) (*Post, error) {
	var err error
	err = db.Debug().Model(&Post{}).Create(&p).Error
	if err != nil {
		return &Post{}, err
	}

	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
		if err != nil {
			return &Post{}, err
		}
	}

	return p, nil
}

func (p *Post) FindAllPosts(db *gorm.DB) (*[]Post, error) {
	var err error
	post := []Post{}

	err = db.Debug().Model(&Post{}).Limit(100).Order("created_at desc").Find(&post).Error
	if err != nil {
		return &[]Post{}, err
	}

	if len(post) > 0 {
		for i, _ := range post {
			err := db.Debug().Model(&User{}).Where("id = ?", post[i].AuthorID).Take(&post[i].Author).Error
			if err != nil {
				return &[]Post{}, err
			}
		}
	}

	return &post, nil
}
