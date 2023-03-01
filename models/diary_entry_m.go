package models

import (
	"errors"
)

var DiaryEntries []DiaryEntry

type DiaryEntry struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

/*
GET
*/

func GetDiaryEntryById(id string) (*DiaryEntry, int, error) {
	for i, element := range DiaryEntries {
		if element.ID == id {
			return &DiaryEntries[i], i, nil
		}
	}
	return nil, -1, errors.New(NOT_FOUND)
}

/*
POST
*/

func PostDiaryEntry(newDiaryEntry DiaryEntry) {
	DiaryEntries = append(DiaryEntries, newDiaryEntry)
	WriteJson()
}

/*
PATCH
*/

func PatchDiaryEntry(field *bool, newValue bool) {
	*field = newValue
	WriteJson()
}

/*
DELETE
*/

func DeleteDiaryEntry(index int) {
	DiaryEntries = Remove(DiaryEntries, index)
	WriteJson()
}
