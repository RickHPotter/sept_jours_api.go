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

// func ToggleDiaryEntryStatus(context *gin.Context) {
// 	id := context.Param("id")
// 	diaryEntry, _, err := models.GetDiaryEntryById(id)
// 	if err != nil {
// 		context.IndentedJSON(http.StatusNotFound, gin.H{"message": NOT_FOUND})
// 		return
// 	}

// 	models.PatchDiaryEntry(&(diaryEntry.Completed), !(diaryEntry.Completed))

// 	context.IndentedJSON(http.StatusOK, diaryEntry)
// }

/*
DELETE
*/

func DeleteDiaryEntry(context *gin.Context) {
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

	// if !DiaryEntry.Completed {
	// 	context.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": UNCOMPLETED_ENTRY})
	// 	return
	// }

	models.DeleteDiaryEntry(index)

	context.IndentedJSON(http.StatusOK, gin.H{"message": DIARY_DELETED})
}
