package controllers

import (
	"net/http"

	"github.com/RickHPotter/flutter_rest_api/models"
	"github.com/gin-gonic/gin"
)

/*
GET
*/

func GetDiaryEntries(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, models.DiaryEntries)
}

func GetDiaryEntry(context *gin.Context) {
	id := context.Param("id")
	DiaryEntry, _, err := models.GetDiaryEntryById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": NOT_FOUND})
		return
	}

	context.IndentedJSON(http.StatusOK, DiaryEntry)
}

/*
POST
*/

func AddDiaryEntry(context *gin.Context) {

	var newDiaryEntry models.DiaryEntry

	err := context.BindJSON(&newDiaryEntry)
	if err != nil {
		return
	}

	models.PostDiaryEntry(newDiaryEntry)

	context.IndentedJSON(http.StatusCreated, newDiaryEntry)
}

/*
PATCH
*/

func UpdateDiaryEntry(context *gin.Context) {
	id, ok := context.GetQuery("id")
	if !ok {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": MISSING_ID})
		return
	}

	var updatedDiaryEntry models.DiaryEntry

	err := context.BindJSON(&updatedDiaryEntry)
	if err != nil {
		return
	}

	if id != updatedDiaryEntry.ID {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": CONFLICTING_ID})
		return
	}

	diaryEntry, _, er := models.GetDiaryEntryById(id)
	if er != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": NOT_FOUND})
		return
	}

	sliceFields := []*string{
		&(diaryEntry.UpdatedAt),
		&(diaryEntry.Content),
	}

	sliceValues := []string{
		updatedDiaryEntry.UpdatedAt,
		updatedDiaryEntry.Content,
	}

	models.PatchDiaryEntryString(sliceFields, sliceValues)
}

/*
DELETE
*/

func DeleteDiaryEntryById(context *gin.Context) {
	id, ok := context.GetQuery("id")
	if !ok {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": MISSING_ID})
		return
	}

	_, index, err := models.GetDiaryEntryById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": NOT_FOUND})
		return
	}

	models.DeleteDiaryEntry(index)

	context.IndentedJSON(http.StatusOK, gin.H{"message": DIARY_DELETED})
}
