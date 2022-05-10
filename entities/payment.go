package entities

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	UserID     uint   `json:"user_id" form:"user_id"`
	VenueID    uint   `json:"venue_id" form:"venue_id"`
	TotalPrice uint   `json:"total_price" form:"total_price"`
	Day        string `json:"day" form:"day"`
	Status     string `json:"status" form:"status"`
	PaymentURL string `json:"payment_url" form:"payment_url"`
	StartDate  string `json:"start_date" form:"start_date"`
	EndDate    string `json:"end_date" form:"end_date"`
	Venue      Venue  `json:"venue" form:"venue"`
	User       User   `json:"user" form:"user"`
}

type TransactionNotificationInput struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}

type PaymentResponse struct {
	ID         uint         `gorm:"primarykey" json:"id"`
	TotalPrice uint         `json:"total_price" form:"total_price"`
	Day        string       `json:"day" form:"day"`
	Status     string       `json:"status" form:"status"`
	PaymentURL string       `json:"payment_url" form:"payment_url"`
	StartDate  string       `json:"start_date" form:"start_date"`
	EndDate    string       `json:"end_date" form:"end_date"`
	User       UserResponse `json:"user" form:"user"`
}
