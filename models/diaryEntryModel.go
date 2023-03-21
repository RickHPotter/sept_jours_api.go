package models

import (
	"time"
)

var DiaryEntries []DiaryEntry

type DiaryEntry struct {
	Hash      string    `gorm:"primarykey" json:"hash"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `gorm:"index"`
}

/*
GET
*/

func GetDiaryEntries() (*[]DiaryEntry, error) {
	var diaries []DiaryEntry
	var err = DB.Find(&diaries).Error

	return &diaries, err
}

func GetDiaryEntryByHash(hash string) (*DiaryEntry, error) {
	diary := DiaryEntry{Hash: hash}
	var err = DB.First(&diary).Error

	return &diary, err
}

/*
POST
*/

func PostDiaryEntry(newDiaryEntry DiaryEntry) error {
	err := DB.Create(&newDiaryEntry).Error
	return err
}

/*
PATCH
*/

func PatchDiaryEntry(diary DiaryEntry) (int, error) {
	db := DB.UpdateColumns(&diary)
	return int(db.RowsAffected), db.Error
}

/*
DELETE
*/

func DeleteDiaryEntry(diary DiaryEntry) (int, error) {
	db := DB.Delete(&diary)
	return int(db.RowsAffected), db.Error
}
