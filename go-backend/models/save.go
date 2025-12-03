package models

import "gorm.io/gorm"

type GameSave struct {
	gorm.Model
	UserID   uint   `json:"user_id"`
	GameName string `json:"game_name"`
	Platform string `json:"platform"`
	Filename string `json:"filename"` // Path to the stored file
	Size     int64  `json:"size"`     // Size in bytes
}
