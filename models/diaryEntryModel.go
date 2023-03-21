package models

import (
	"errors"
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

func GetDiaryEntryByHash(hash string) (*DiaryEntry, int, error) {
	for i, element := range DiaryEntries {
		if element.Hash == hash {
			return &DiaryEntries[i], i, nil
		}
	}
	return nil, -1, errors.New(NOT_FOUND)
}

/*
POST
*/

func PostDiaryEntry(newDiaryEntry DiaryEntry) bool {
	DB.Create(&newDiaryEntry)

	return newDiaryEntry.Hash != ""
	// soon to be depracated
	// DiaryEntries = append(DiaryEntries, newDiaryEntry)
	// WriteJson()
}

/*
PATCH
*/

func PatchDiaryEntry(diary *DiaryEntry, updated DiaryEntry) {
	diary.Title = updated.Title
	diary.Content = updated.Content
	diary.UpdatedAt = updated.UpdatedAt

	WriteJson()
}

/*
DELETE
*/

func DeleteDiaryEntry(diary DiaryEntry) bool {
	DB.Delete(&diary)
	return true // TODO: HOW CAN IT BE FALSE THOUGH

	// soon to be depracated
	// DiaryEntries = Remove(DiaryEntries, index)
	// WriteJson()
}
