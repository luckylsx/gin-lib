package models

import "time"

// Model base model
type Model struct {
	ID        int       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDeleted int       `json:"is_deleted"`
}

// ConVo database connect info
type ConVo struct {
	Driver  string
	SQLType string
}
