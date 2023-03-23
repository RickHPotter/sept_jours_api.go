package models

import (
	"gorm.io/gorm"
)

type RequestBody struct {
	Email    string
	Password string
}

type User struct {
	gorm.Model        // ID `gorm:"primarykey"`, createdAt, updatedAt, deletedAt `gorm:"index"`
	Email      string `gorm:"unique"`
	Password   string
}

/*
! GET
*/

func GetUser(cond string, condVar any) (*User, error) {
	var user = User{}
	db := DB.First(&user, cond, condVar)
	return &user, db.Error
}

/*
! POST
*/

func CreateUser(user User) error {
	return DB.Create(&user).Error
}
