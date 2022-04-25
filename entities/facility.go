package entities

import (
	"gorm.io/gorm"
)

type Facility struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	IconName string `json:"icon_name" form:"icon_name"`
}
