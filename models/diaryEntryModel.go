package models

import (
	"time"

	"gorm.io/gorm"
)

var DiaryEntries []DiaryEntry

type DiaryEntry struct {
	Hash      string    `gorm:"primarykey" json:"hash"`
	UserId    int       `gorm:"not null"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `gorm:"index"`
}

/*
! GET
*/

func GetDiaryEntries(userId int) (*[]DiaryEntry, error) {
	var diaries []DiaryEntry
	var err = DB.Find(&diaries, "user_id = ?", userId).Error

	return &diaries, err
}

func GetDiaryEntryByHash(hash string) (*DiaryEntry, error) {
	diary := DiaryEntry{Hash: hash}
	var err = DB.First(&diary).Error

	return &diary, err
}

/*
! POST
*/

func PostDiaryEntry(newDiaryEntry DiaryEntry) error {
	err := DB.Create(&newDiaryEntry).Error
	return err
}

/*
! PATCH
*/

func PatchDiaryEntry(diary DiaryEntry) (int, error) {
	var db *gorm.DB

	// i know this is horrendous, but no way I found to make time.Time act like the atrocity postgress timestamp looks like
	if diary.UpdatedAt == (time.Time{}) {
		db = DB.Model(&diary).Updates(DiaryEntry{Title: diary.Title, Content: diary.Content})
	} else {
		db = DB.UpdateColumns(&diary)
	}

	return int(db.RowsAffected), db.Error
}

/*
! DELETE
*/

func DeleteDiaryEntry(diary DiaryEntry) (int, error) {
	db := DB.Delete(&diary)
	return int(db.RowsAffected), db.Error
}
