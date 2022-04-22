package entities

import (
	"gorm.io/gorm"
)

type PaymentLink struct {
	gorm.Model
	SecretKey      string     `json:"secret_key" form:"secret_key"`
	Code           uint       `json:"code" form:"code"`
}