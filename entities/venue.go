package entities

import (
	"time"

	"gorm.io/gorm"
)

type Venue struct {
	gorm.Model
	UserID        uint        `json:"user_id" form:"user_id"`
	OperationalID uint        `json:"operational_id" form:"operational_id"`
	FacilityID    []uint      `json:"facility_id" form:"facility_id"`
	CategoryID    uint        `json:"category_id" form:"category_id"`
	Image         string      `json:"image" form:"image"`
	Name          string      `json:"name" form:"name"`
	City          string      `json:"city" form:"city"`
	Address       string      `json:"address" form:"address"`
	BookingTime   time.Time   `json:"booking_time" form:"booking_time"`
	Operational   Operational `gorm:"foreignKey:OperationalID;references:ID" json:"operational" form:"operational"`
	User          User        `gorm:"foreignKey:UserID;references:ID" json:"user" form:"user"`
	Facility      Facility    `gorm:"foreignKey:FacilityID;references:ID" json:"facility" form:"facility"`
	Category      Category    `gorm:"foreignKey:CategoryID;references:ID" json:"category" form:"category"`
}
