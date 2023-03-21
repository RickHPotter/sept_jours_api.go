package controllers

import (
	"net/http"

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
