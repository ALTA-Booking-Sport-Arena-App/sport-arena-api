package entities

import (
	"gorm.io/gorm"
)

type Facility struct {
	gorm.Model
	Name           string       `json:"name" form:"name"`
}