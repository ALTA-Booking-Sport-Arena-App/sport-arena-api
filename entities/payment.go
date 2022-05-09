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
	Venue      Venue
	User       User
}

type TransactionNotificationInput struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}
