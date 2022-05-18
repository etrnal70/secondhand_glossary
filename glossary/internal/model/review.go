package model

import "time"

type Review struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	User      User      `json:"user"`
	DeviceID  uint      `json:"device_id"`
	Device    Device    `json:"device"`
	Score     uint      `gorm:"not null" json:"score"`
	Post      string    `gorm:"not null" json:"post"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
