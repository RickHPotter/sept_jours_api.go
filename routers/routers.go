package routers

import (
	"github.com/RickHPotter/flutter_web_api/controllers"
	"github.com/RickHPotter/flutter_web_api/models"
	"github.com/gin-gonic/gin"
)

func LoadRouters() {
	models.ReadJson()

	router := gin.Default()

	// GET
	router.GET("/diary/entry", controllers.GetDiaryEntries)
	router.GET("/diary/entry/:id", controllers.GetDiaryEntry)

	// POST
	router.POST("/diary/entry", controllers.AddDiaryEntry)

	// PATCH
	router.PATCH("/diary/entry/:id", controllers.ToggleDiaryEntryStatus)

	// DELETE
	router.DELETE("/diary/entry/delete/", controllers.DeleteDiaryEntry)

	/* RUNNING */
	router.Run("192.168.100.66:9090")
}
