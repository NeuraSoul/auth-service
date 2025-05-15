package models 

import (
	"time"
)


type Users struct {
	ID int `gorm:"increment"`
	UserID string 
	Email string
	Password string
	Created time.Time
}


type AuthToken struct {
	ID int
	Token string
	ExpiresAT time.Time
}

