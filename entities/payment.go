package entities

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	UserID     int       `json:"user_id" form:"user_id"`
	VenueID    int       `json:"venue_id" form:"venue_id"`
	TotalPrice int       `json:"total_price" form:"total_price"`
	Status     string    `json:"status" form:"status"`
	StartDate  time.Time `json:"start_date" form:"start_date"`
	EndDate    time.Time `json:"end_date" form:"end_date"`
	User       User      `gorm:"foreignKey:UserID;references:ID" json:"user" form:"user"`
	Venue      Venue     `gorm:"foreignKey:VenueID;references:ID" json:"venue" form:"venue"`
}

// type PaymentResponse struct {
// 	ID     uint          `gorm:"primaryKey" json:"id" form:"id"`
// 	Status string        `json:"status" form:"status"`
// 	Venue  VenueResponse `json:"venue" form:"venue"`
// }
