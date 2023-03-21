package controllers

import (
	"net/http"

	"github.com/RickHPotter/flutter_rest_api/models"
	"github.com/gin-gonic/gin"
)

/*
! GET
*/

func GetDiaryEntries(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, models.DiaryEntries)
}

func GetDiaryEntry(context *gin.Context) {
	id := context.Param("hash")
	DiaryEntry, _, err := models.GetDiaryEntryByHash(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": NOT_FOUND})
		return
	}

	context.IndentedJSON(http.StatusOK, DiaryEntry)
}

/*
! POST
*/

func AddDiaryEntry(context *gin.Context) {
	// ! Get Diary Entry off Request Body
	var newDiaryEntry models.DiaryEntry

	err := context.ShouldBindJSON(&newDiaryEntry)
	if err != nil {
		BadReq(context, "Something wrong with the Request.\n"+err.Error())
		return
	}

	// ! Check if given Diary Entry Hash is in the Database already
	var checkDiaryEntry models.DiaryEntry
	// TODO: initialisers.DB.First(&checkDiaryEntry, "hash = ?", newDiaryEntry.Hash)

	if checkDiaryEntry.Hash != "" {
		BadReq(context, "Diary Entry already exists. No Duplicates Allowed.")
		return
	}

	// ! Create Diary Entry
	result := models.PostDiaryEntry(newDiaryEntry)

	if !result {
		BadReq(context, "Failed to create Diary Entry in the Database.")
		return
	}

	context.IndentedJSON(http.StatusCreated, newDiaryEntry)
}

/*
! PATCH
*/

func UpdateDiaryEntry(context *gin.Context) {
	hash, ok := context.GetQuery("hash")
	if !ok {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": MISSING_ID})
		return
	}

	var updatedDiaryEntry models.DiaryEntry

	err := context.BindJSON(&updatedDiaryEntry)
	if err != nil {
		return
	}

	if hash != updatedDiaryEntry.Hash {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": CONFLICTING_ID})
		return
	}

	diaryEntry, _, er := models.GetDiaryEntryByHash(hash)
	if er != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": NOT_FOUND})
		return
	}

	models.PatchDiaryEntry(diaryEntry, updatedDiaryEntry)
}

/*
! DELETE
*/

func DeleteDiaryEntryByHash(context *gin.Context) {
	// ! Get Diary Entry Hash off Req Body
	hash, ok := context.GetQuery("hash")
	if !ok {
		BadReq(context, MISSING_ID)
		return
	}

	// TODO: CHECK IF TO-BE-DELETED ENTRY IS ACTUALLY THERE
	// _, index, err := models.GetDiaryEntryByHash(hash)
	// if err != nil {
	// 	NotFound(context, NOT_FOUND)
	// 	return
	// }

	// ! Delete Diary Entry with such Hash
	var diary models.DiaryEntry
	diary.Hash = hash
	result := models.DeleteDiaryEntry(diary)

	if !result {
		BadReq(context, "Failed to delete Diary Entry from the Database.")
		return
	}

	Ok(context, DIARY_DELETED)
}
