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
	router.GET("/api/v1/diary", middleware.RequireAuth, controllers.GetDiaryEntries)
	router.GET("/api/v1/diary/:hash", middleware.RequireAuth, controllers.GetDiaryEntry)

	// * POST
	router.POST("/api/v1/diary/insert", middleware.RequireAuth, controllers.AddDiaryEntry)

	// * PATCH
	router.PATCH("/api/v1/diary/update", middleware.RequireAuth, controllers.UpdateDiaryEntry)

	// * DELETE
	router.DELETE("/api/v1/diary/delete", middleware.RequireAuth, controllers.DeleteDiaryEntryByHash)

	// ! USERS

	// * GET
	router.GET("/logout", controllers.Logout)

	// * POST
	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.Login)

	// ? RUNNING
	router.Run()

	// ? SERVING
	http.ListenAndServe(os.Getenv("PORT"), nil)
}
