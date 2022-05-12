package entities

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name     string  `json:"name" form:"name"`
	IconName string  `json:"icon_name" form:"icon_name"`
	Venue    []Venue `gorm:"foreignKey:CategoryID;references:ID"`
}

type CategoryResponse struct {
	ID       uint   `gorm:"primarykey" json:"id"`
	Name     string `json:"name" form:"name"`
	IconName string `json:"icon_name" form:"icon_name"`
}
