package model

import "time"

type Scores struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	DeviceID     uint      `gorm:"not null" json:"device_id"`
	UserScore    uint      `gorm:"default:0" json:"user_score"`
	CrawlerScore uint      `gorm:"default:0" json:"crawler_score"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
