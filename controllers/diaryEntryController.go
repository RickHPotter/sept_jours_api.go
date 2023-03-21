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
	if diaryEntries, err := models.GetDiaryEntries(); err != nil {
		BadReq(context, "What could've gone wrong?")
	} else {
		context.IndentedJSON(http.StatusOK, diaryEntries)
	}
}

func GetDiaryEntry(context *gin.Context) {
	hash := context.Param("hash")
	diaryEntry, err := models.GetDiaryEntryByHash(hash)
	if err != nil {
		NotFound(context, NOT_FOUND)
		return
	}

	context.IndentedJSON(http.StatusOK, diaryEntry)
}

/*
! POST
*/

func AddDiaryEntry(context *gin.Context) {
	// ! Get Diary Entry off Request Body
	var newDiaryEntry models.DiaryEntry

	if err := context.ShouldBindJSON(&newDiaryEntry); err != nil {
		BadReq(context, "Something wrong with the Request.\n"+err.Error())
		return
	}

	// ! Check if given Diary Entry Hash is in the Database already
	if _, err := models.GetDiaryEntryByHash(newDiaryEntry.Hash); err == nil {
		BadReq(context, "Diary Entry already exists. No Duplicates Allowed.")
		return
	}

	// ! Create Diary Entry
	if err := models.PostDiaryEntry(newDiaryEntry); err != nil {
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
		BadReq(context, MISSING_ID)
		return
	}

	var diaryEntry models.DiaryEntry

	err := context.BindJSON(&diaryEntry)
	if err != nil {
		BadReq(context, "Failed to fetch request.")
		return
	}

	if hash != diaryEntry.Hash {
		BadReq(context, CONFLICTING_ID)
		return
	}

	rowsAffected, err := models.PatchDiaryEntry(diaryEntry)
	if rowsAffected == 0 {
		NotFound(context, NOT_FOUND)
		return
	}
	if err != nil {
		BadReq(context, "Failed to update Diary Entry in the Database.")
		return
	}

	context.IndentedJSON(http.StatusOK, diaryEntry)
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

	// ! Delete Diary Entry with such Hash
	var diary models.DiaryEntry
	diary.Hash = hash

	rowsAffected, err := models.DeleteDiaryEntry(diary)

	if rowsAffected == 0 {
		NotFound(context, NOT_FOUND)
		return
	}
	if err != nil {
		BadReq(context, "Failed to delete Diary Entry from the Database.")
		return
	}

	Ok(context, DIARY_DELETED)
}
