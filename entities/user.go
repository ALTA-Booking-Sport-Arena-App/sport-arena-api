package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Fullname      string `gorm:"not null" json:"fullname" form:"fullname"`
	Username      string `gorm:"not null;unique" json:"username" form:"username"`
	Email         string `gorm:"not null;unique" json:"email" form:"email"`
	Password      string `gorm:"not null" json:"password" form:"password"`
	PhoneNumber   string `gorm:"not null" json:"phone_number" form:"phone_number"`
	Role          string `gorm:"not null" json:"role" form:"role"`
	Image         string `json:"image" form:"image"`
	BusinessName  string `json:"business_name" form:"business_name"`
	BusinessType  string `json:"business_type" form:"business_type"`
	BusinessTerms string `json:"business_terms" form:"business_terms"`
}

type UserResponse struct {
	gorm.Model
	FullName      string `gorm:"not null" json:"fullname" form:"fullname"`
	Username      string `gorm:"not null;unique" json:"username" form:"username"`
	Email         string `gorm:"not null;unique" json:"email" form:"email"`
	PhoneNumber   string `gorm:"not null" json:"phone_number" form:"phone_number"`
	Role          string `gorm:"not null" json:"role" form:"role"`
	Image         string `json:"image" form:"image"`
	BusinessName  string `json:"business_name" form:"business_name"`
	BusinessType  string `json:"business_type" form:"business_type"`
	BusinessTerms string `json:"business_terms" form:"business_terms"`
}
