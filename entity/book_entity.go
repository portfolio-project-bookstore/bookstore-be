package entity

import "time"

type Book struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Category    string
	Author      string
	Description string
	Price       int
	Stock       int
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	UsersID     uint      `gorm:"default:NULL"`
}
