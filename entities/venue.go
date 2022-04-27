package entities

import (
	"time"

	"gorm.io/gorm"
)

type Venue struct {
	gorm.Model
	UserID        uint            `json:"user_id" form:"user_id"`
	CategoryID    uint            `json:"category_id" form:"category_id"`
	Image         string          `json:"image" form:"image"`
	Name          string          `json:"name" form:"name"`
	City          string          `json:"city" form:"city"`
	Address       string          `json:"address" form:"address"`
	BookingTime   time.Time       `json:"booking_time" form:"booking_time"`
	Step2         []Step2         `gorm:"foreignKey:VenueID;references:ID" json:"operational" form:"operational"`
	User          User            `gorm:"foreignKey:UserID;references:ID" json:"user" form:"user"`
	Category      Category        `gorm:"foreignKey:CategoryID;references:ID" json:"category" form:"category"`
	VenueFacility []VenueFacility `gorm:"foreignKey:VenueID;references:ID" json:"facility_id" form:"facility_id"`
}

type GetVenuesResponse struct {
	ID       uint   `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Location string `json:"location" form:"location"`
	Image    string `json:"image" form:"image"`
}
