package models

import "gorm.io/gorm"

type RequestBody struct {
	Email    string
	Password string
}

type User struct {
	gorm.Model        // ID `gorm:"primarykey"`, createdAt, updatedAt, deletedAt `gorm:"index"`
	Email      string `gorm:"unique"`
	Password   string
}
