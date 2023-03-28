package controllers

import (
	"fmt"
	"net/http"

	"github.com/RickHPotter/flutter_rest_api/models"
	"github.com/gin-gonic/gin"
)

/*
! GET
*/

func GetDiaryEntries(context *gin.Context) {
	// ! Get user from cookie to check which user is using this API call
	userCookie, exists := context.Get("user")
	if !exists {
		BadReq(context, "Failed to fetch user variable from middleware.")
		return
	}

	user := userCookie.(models.User)

	// ! Retrieve Diary Entries from given user_id
	if diaryEntries, err := models.GetDiaryEntries(int(user.ID)); err != nil {
		fmt.Print(err.Error())
		BadReq(context, "What could've gone wrong?")
	} else {
		context.IndentedJSON(http.StatusOK, diaryEntries)
	}
}

func GetDiaryEntry(context *gin.Context) {
	// ! Get hash off req body
	hash := context.Param("hash")
	diaryEntry, err := models.GetDiaryEntryByHash(hash)
	if err != nil {
		NotFound(context, NOT_FOUND)
		return
	}

	// ! Check if isAuthor
	if isNotAuthor(context, *diaryEntry) {
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

	// ! Check if userId was informed
	if newDiaryEntry.UserId == 0 {
		BadReq(context, "Orphan Diary Entries not allowed. Get a UserId.")
		return
	}

	// ! Check if userId of JSON is the same as userId of Cookie
	if isNotAuthor(context, newDiaryEntry) {
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
	// ! Get hash off req body
	hash, ok := context.GetQuery("hash")
	if !ok {
		BadReq(context, MISSING_ID)
		return
	}

	var diaryEntry models.DiaryEntry

	if err := context.ShouldBindJSON(&diaryEntry); err != nil {
		BadReq(context, "Something wrong with the Request.\n"+err.Error())
		return
	}

	// ! Check if query hash is the same as JSON hash
	if hash != diaryEntry.Hash {
		BadReq(context, CONFLICTING_ID)
		return
	}

	// ! Check if isAuthor
	if isNotAuthor(context, diaryEntry) {
		return
	}

	// ! Perform update and validate if it did happen
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

	// ! Check if Diary Exists
	diaryEntry, err := models.GetDiaryEntryByHash(hash)
	if err != nil {
		BadReq(context, NOT_FOUND)
		return
	}

	// ! Check if isAuthor
	if isNotAuthor(context, *diaryEntry) {
		return
	}

	// ! Delete Diary Entry with such Hash
	var diary models.DiaryEntry
	diary.Hash = hash

	rowsAffected, err := models.DeleteDiaryEntry(diary)

	if rowsAffected == 0 {
		NotFound(context, NOT_FOUND+" .DELETE. ")
		return
	}
	if err != nil {
		BadReq(context, "Failed to delete Diary Entry from the Database.")
		return
	}

	Ok(context, DIARY_DELETED)
}

func Ping(c *gin.Context) {
	c.IndentedJSON(http.StatusAccepted, gin.H{
		"Message": "Pong.",
	})
}
