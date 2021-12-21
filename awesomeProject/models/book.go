package models

import (
	"time"
)

type Books struct {
	ID          uint64         `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" binding:"required"`
	Description string         `json:"description"`
	MediumPrice float32         `json:"medium_price"`
	Author      string         `json:"author" binding:"required"`
	ImageURL    string         `json:"img_url"`
	CreatedAt   time.Time      `json:"created"`
	UpdatedAt   time.Time      `json:"updated"`
	DeletedAt   *time.Time `gorm:"index" json:"deleted"`
}
type BookTeste struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" binding:"required"`
	Description string         `json:"description"`
	MediumPrice float32         `json:"medium_price"`
	Author      string         `json:"author" binding:"required"`
	ImageURL    string         `json:"img_url"`
	CreatedAt   time.Time      `json:"created"`
	UpdatedAt   time.Time      `json:"updated"`
	DeletedAt   *time.Time `gorm:"index" json:"deleted"`
}
