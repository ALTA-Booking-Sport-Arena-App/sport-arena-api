package entities

import (
	"time"

	"gorm.io/gorm"
)

type Operational struct {
	gorm.Model
	Day         []string      `json:"day" form:"day"`
	OpenHour    time.Time     `json:"open_hour" form:"open_hour"`
	CloseHour   time.Time     `json:"close_hour" form:"close_hour"`
	Price        uint         `json:"price" form:"price"`
}