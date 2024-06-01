package models

import (
	"time"
)

type User struct {
	ID uint `gorm:"primary_key" json:"id"`
	Username string `gorm:"type:varchar(100);not null" json:"username"`
	Email string `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password string `gorm:"type:varchar(100);min:6;not null" json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Photos []Photo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}