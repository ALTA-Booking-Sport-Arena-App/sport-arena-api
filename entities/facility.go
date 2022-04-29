package entities

import (
	"gorm.io/gorm"
)

type Facility struct {
	gorm.Model
	Name          string          `json:"name" form:"name"`
	IconName      string          `json:"icon_name" form:"icon_name"`
	VenueFacility []VenueFacility `gorm:"foreignKey:FacilityID;references:ID" json:"facility_id" form:"facility_id"`
}
