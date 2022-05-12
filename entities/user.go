package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName            string  `gorm:"not null" json:"fullname" form:"fullname"`
	Username            string  `gorm:"not null;unique" json:"username" form:"username"`
	Email               string  `gorm:"not null;unique" json:"email" form:"email"`
	Password            string  `gorm:"not null" json:"password" form:"password"`
	PhoneNumber         string  `gorm:"not null" json:"phone_number" form:"phone_number"`
	Role                string  `gorm:"default:user" json:"role" form:"role"`
	Image               string  `json:"image" form:"image"`
	BusinessName        string  `json:"business_name" form:"business_name"`
	BusinessDescription string  `json:"business_description" form:"business_description"`
	BusinessCertificate string  `json:"business_certificate" form:"business_certificate"`
	Status              string  `json:"status" form:"status"`
	Venues              []Venue `gorm:"foreignKey:UserID;references:ID" json:"venue" form:"venue"`
}

type UserResponse struct {
	gorm.Model
	FullName            string `gorm:"not null" json:"fullname" form:"fullname"`
	Username            string `gorm:"not null;unique" json:"username" form:"username"`
	Email               string `gorm:"not null;unique" json:"email" form:"email"`
	PhoneNumber         string `gorm:"not null" json:"phone_number" form:"phone_number"`
	Role                string `gorm:"not null" json:"role" form:"role"`
	Image               string `json:"image" form:"image"`
	BusinessName        string `json:"business_name" form:"business_name"`
	BusinessDescription string `json:"business_description" form:"business_description"`
	BusinessCertificate string `json:"business_certificate" form:"business_certificate"`
	Status              string `json:"status" form:"status"`
}

type ListUsersResponse struct {
	ID          uint   `gorm:"primarykey" json:"id"`
	FullName    string `gorm:"not null" json:"fullname" form:"fullname"`
	Username    string `gorm:"not null;unique" json:"username" form:"username"`
	Email       string `gorm:"not null;unique" json:"email" form:"email"`
	PhoneNumber string `gorm:"not null" json:"phone_number" form:"phone_number"`
	Image       string `json:"image" form:"image"`
}

type ListOwnersResponse struct {
	gorm.Model
	FullName     string  `gorm:"not null" json:"fullname" form:"fullname"`
	Username     string  `gorm:"not null;unique" json:"username" form:"username"`
	Image        string  `json:"image" form:"image"`
	BusinessName string  `json:"business_name" form:"business_name"`
	Venues       []Venue `gorm:"foreignKey:UserID;references:ID" json:"venue" form:"venue"`
}

type ListOwnerRequestResponse struct {
	gorm.Model
	FullName            string `gorm:"not null" json:"fullname" form:"fullname"`
	Username            string `gorm:"not null;unique" json:"username" form:"username"`
	Email               string `gorm:"not null;unique" json:"email" form:"email"`
	PhoneNumber         string `gorm:"not null" json:"phone_number" form:"phone_number"`
	Status              string `json:"status" form:"status"`
	BusinessCertificate string `json:"business_certificate" form:"business_certificate"`
}
