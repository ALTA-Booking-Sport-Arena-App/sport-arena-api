package entities

import (
	"gorm.io/gorm"
)

type VenueFacility struct {
	gorm.Model
	VenueID    uint     `json:"venue_id" form:"venue_id"`
	FacilityID uint     `json:"facility_id" form:"facility_id"`
	Facility   Facility `gorm:"foreignKey:FacilityID;references:ID" json:"facility" form:"facility"`
}
