package routers

import (
	"github.com/RickHPotter/flutter_rest_api/controllers"
	"github.com/RickHPotter/flutter_rest_api/models"
	"github.com/gin-gonic/gin"
)

func LoadRouters() {
	models.ReadJson()

	router := gin.Default()

	// GET
	router.GET("/diary/entry", controllers.GetDiaryEntries)
	router.GET("/diary/entry/:hash", controllers.GetDiaryEntry)

	// POST
	router.POST("/diary/entry/insert", controllers.AddDiaryEntry)

	// PATCH
	router.PATCH("/diary/entry/update", controllers.UpdateDiaryEntry)

	// DELETE
	router.DELETE("/diary/entry/delete", controllers.DeleteDiaryEntryByHash)

	/* RUNNING */
	router.Run(":9090")
}
