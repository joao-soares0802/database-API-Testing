package models

import "time"

type Games struct {
	ID        uint64     `json:"id" gorm:"primaryKey"`
	Title     string     `json:"title" binding:"required"`
	Genre     string     `json:"genre"`
	Publisher string     `json:"publisher"`
	Price     float32    `json:"price"`
	CreatedAt time.Time  `json:"created"`
	UpdatedAt time.Time  `json:"updated"`
	DeletedAt *time.Time `gorm:"index" json:"deleted"`
}

func (Games) TableName() string {
	return "games"
}
