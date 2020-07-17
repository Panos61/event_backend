package models

import "github.com/jinzhu/gorm"

//CityEvents => Find Events based on City Param.
func (e *Events) CityEvents(db *gorm.DB, city string) (*[]Events, error) {
	var err error
	events := []Events{}

	err = db.Debug().Model(&Events{}).Where("city = ?", city).Limit(150).Order("created_at desc").Find(&events).Error
	if err != nil {
		return &[]Events{}, err
	}

	if len(events) > 0 {
		for i, _ := range events {
			err := db.Debug().Model(&Events{}).Where("city = ?", events[i].City).Take(&events[i]).Error
			if err != nil {
				return &[]Events{}, err
			}
		}
	}

	return &events, err
}
