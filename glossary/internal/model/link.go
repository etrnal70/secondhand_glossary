package model

import "time"

type Link struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	DeviceID  uint      `json:"device_id"`
	Store     string    `json:"store"`
	Link      string    `gorm:"not null" json:"link"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
