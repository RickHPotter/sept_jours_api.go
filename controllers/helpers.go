package controllers

import (
	"net/http"

	"github.com/RickHPotter/flutter_rest_api/models"
	"github.com/gin-gonic/gin"
)

const NOT_FOUND = "DIARY entry not found."
const MISSING_ID = "Missing DIARY HASH query parameter."
const CONFLICTING_ID = "JSON HASH other than the URL HASH Query."
const DIARY_DELETED = "DIARY entry deleted."

/*
! HELPERS
*/

func BadReq(c *gin.Context, errorMessage string) {
	c.IndentedJSON(http.StatusBadRequest, gin.H{
		"Error": errorMessage,
	})
}

func Unauthorised(c *gin.Context, errorMessage string) {
	c.IndentedJSON(http.StatusUnauthorized, gin.H{
		"Error": errorMessage,
	})
}

func NotFound(context *gin.Context, errorMessage string) {
	context.IndentedJSON(http.StatusNotFound, gin.H{
		"Error": errorMessage,
	})
}

func Ok(context *gin.Context, message string) {
	context.IndentedJSON(http.StatusOK, gin.H{
		"Message": message,
	})
}

func isAuthor(context *gin.Context, diaryEntry models.DiaryEntry) bool {
	userCookie, exists := context.Get("user")
	if !exists {
		BadReq(context, "Failed to fetch user variable from middleware.")
		return false
	}

	user := userCookie.(models.User)

	if int(user.ID) != diaryEntry.UserId {
		Unauthorised(context, "Keep your hands to yourself, little one.")
		return false
	}
	return true
}

func isNotAuthor(context *gin.Context, diaryEntry models.DiaryEntry) bool {
	return !isAuthor(context, diaryEntry)
}
