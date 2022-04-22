package entities

import (
	"gorm.io/gorm"
)

type History struct {
	gorm.Model
	PaymentID uint    `json:"payment_id" form:"payment_id"`
	Status    string  `json:"status" form:"status"`
	Payment   Payment `gorm:"foreignKey:PaymentID;references:ID" json:"payment" form:"payment"`
}
