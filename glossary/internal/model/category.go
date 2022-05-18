package model

import "time"

type Category struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Category  string    `gorm:"not null" json:"category"`
	Devices   []*Device `json:"devices"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
