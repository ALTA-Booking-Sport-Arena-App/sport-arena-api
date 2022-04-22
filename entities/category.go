package entities

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `json:"name" form:"name"`
}
