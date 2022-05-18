package model

import "time"

type Trait struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Trait     string    `gorm:"not null" json:"trait"`
	Context   string    `gorm:"not null" json:"context"`
	Devices   []*Device  `gorm:"many2many:device_traits;" json:"devices"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
