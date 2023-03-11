package models

import (
	"errors"
)

var DiaryEntries []DiaryEntry

type DiaryEntry struct {
	Hash      string `json:"hash"`
	Title     string `json:"title"`
	Content   string `json:"content,"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
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

func PostDiaryEntry(newDiaryEntry DiaryEntry) {
	DiaryEntries = append(DiaryEntries, newDiaryEntry)
	WriteJson()
}

/*
PATCH
*/

func PatchDiaryEntryBoolean(field *bool, newValue bool) {
	*field = newValue
	WriteJson()
}

func PatchDiaryEntryString(field []*string, newValue []string) {
	for i := range field {
		*field[i] = newValue[i]
	}
	WriteJson()
}

/*
DELETE
*/

func DeleteDiaryEntry(index int) {
	DiaryEntries = Remove(DiaryEntries, index)
	WriteJson()
}
