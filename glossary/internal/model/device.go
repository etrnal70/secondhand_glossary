package model

import (
	"database/sql"
	"time"
)

type Device struct {
	ID           uint         `gorm:"primaryKey" json:"id"`
	Traits       []*Trait     `gorm:"many2many:device_traits;" json:"traits"`
	Reviews      []*Review    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"reviews"`
	Links        []*Link      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"links"`
	Scores       Scores       `json:"scores"`
	CategoryID   uint         `json:"category_id"`
	Category     Category     `gorm:"foreignKey:CategoryID" json:"category"`
	Manufacturer string       `gorm:"not null" json:"manufacturer"`
	Lineup       string       `json:"lineup"`
	Type         string       `gorm:"not null" json:"type"`
	Image        string       `json:"image"`
  ReleaseDate  sql.NullTime `json:"release_date" swaggertype:"primitive,integer"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}
