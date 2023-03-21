package routers

import (
	"net/http"
	"os"

	"github.com/RickHPotter/flutter_rest_api/controllers"
	"github.com/RickHPotter/flutter_rest_api/middleware"
	"github.com/gin-gonic/gin"
)

func LoadRouters() {
	router := gin.Default()

	// ! DIARY ENTRY

	// * GET
	router.GET("/api/v1/diary", controllers.GetDiaryEntries)
	router.GET("/api/v1/diary/:hash", controllers.GetDiaryEntry)

	// * POST
	router.POST("/api/v1/diary/insert", middleware.RequireAuth, controllers.AddDiaryEntry)

	// * PATCH
	router.PATCH("/api/v1/diary/update", controllers.UpdateDiaryEntry)

	// * DELETE
	router.DELETE("/api/v1/diary/delete", controllers.DeleteDiaryEntryByHash)

	// ! USERS

	// * GET
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)
	router.GET("/logout", controllers.Logout)

	// * POST
	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.Login)

	// ? RUNNING
	router.Run()

	// ? SERVING
	http.ListenAndServe(os.Getenv("PORT"), nil)
}
