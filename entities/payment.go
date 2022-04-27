package entities

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	UserID        uint        `json:"user_id" form:"user_id"`
	VenueID       uint        `json:"venue_id" form:"venue_id"`
	PaymentLinkID uint        `json:"payment_link_id" form:"payment_link_id"`
	TotalPrice    uint        `json:"total_price" form:"total_price"`
	Status        string      `json:"status" form:"status"`
	StartDate     time.Time   `json:"start_date" form:"start_date"`
	EndDate       time.Time   `json:"end_date" form:"end_date"`
	User          User        `gorm:"foreignKey:UserID;references:ID" json:"user" form:"user"`
	Venue         Venue       `gorm:"foreignKey:VenueID;references:ID" json:"venue" form:"venue"`
	PaymentLink   PaymentLink `gorm:"foreignKey:PaymentLinkID;references:ID" json:"payment_link" form:"payment_link"`
}

// type PaymentResponse struct {
// 	ID     uint          `gorm:"primaryKey" json:"id" form:"id"`
// 	Status string        `json:"status" form:"status"`
// 	Venue  VenueResponse `json:"venue" form:"venue"`
// }
