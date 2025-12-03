package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email               string    `gorm:"unique;not null" json:"email"`
	Password            string    `json:"-"`
	SubscriptionStatus  string    `json:"subscription_status" gorm:"default:'inactive'"` // active, inactive
	SubscriptionPlan    string    `json:"subscription_plan" gorm:"default:'none'"`      // monthly, yearly, none
	SubscriptionExpiresAt *time.Time `json:"subscription_expires_at"`
}
