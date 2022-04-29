package entities

import (
	"gorm.io/gorm"
)

type Step2 struct {
	gorm.Model
	VenueID   uint   `json:"venue_id" form:"venue_id"`
	Day       string `json:"day" form:"day"`
	OpenHour  string `json:"open_hour" form:"open_hour"`
	CloseHour string `json:"close_hour" form:"close_hour"`
	Price     uint   `json:"price" form:"price"`
}
