package models

/*
OPERATIONS
*/

func Remove(slice []DiaryEntry, index int) []DiaryEntry {
	if len(slice) > 0 {
		return append(slice[:index], slice[index+1:]...)
	}
	return slice
}

func RemoveAll(slice []DiaryEntry) []DiaryEntry {
	return append(slice[0:0], slice[0:0]...)
}

func Insert(slice []DiaryEntry, newDiaryEntry DiaryEntry, index int) []DiaryEntry {
	if len(slice) >= index {
		slice = append(slice[:index+1], slice[index:]...)
		slice[index] = newDiaryEntry
	}
	return slice
}
