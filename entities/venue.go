package entities

import (
	"gorm.io/gorm"
)

type Venue struct {
	gorm.Model
	UserID        uint            `json:"user_id" form:"user_id"`
	CategoryID    uint            `json:"category_id" form:"category_id"`
	Image         string          `json:"image" form:"image"`
	Name          string          `json:"name" form:"name"`
	Description   string          `json:"description" form:"description"`
	City          string          `json:"city" form:"city"`
	Address       string          `json:"address" form:"address"`
	Step2         []Step2         `gorm:"foreignKey:VenueID;references:ID" json:"operational_hours" form:"operational_hours"`
	User          User            `gorm:"foreignKey:UserID;references:ID" json:"user" form:"user"`
	Category      Category        `gorm:"foreignKey:CategoryID;references:ID" json:"category" form:"category"`
	VenueFacility []VenueFacility `gorm:"foreignKey:VenueID;references:ID" json:"facility_id" form:"facility_id"`
	Payment       []Payment       `json:"payment" form:"payment"`
}

type GetVenuesResponse struct {
	ID       uint   `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Category string `json:"category" form:"category"`
	City     string `json:"location" form:"location"`
	Image    string `json:"image" form:"image"`
	Price    uint   `json:"price" form:"price"`
}

type VenueResponse struct {
	ID          uint              `gorm:"primarykey" json:"id"`
	UserID      uint              `json:"user_id" form:"user_id"`
	Image       string            `json:"image" form:"image"`
	Name        string            `json:"name" form:"name"`
	Description string            `json:"description" form:"description"`
	City        string            `json:"city" form:"city"`
	Address     string            `json:"address" form:"address"`
	Payment     []PaymentResponse `json:"payment" form:"payment"`
}
