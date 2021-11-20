package routes

import (
	. "github.com/bilou4/GoDirBrowser/constants"
	"github.com/bilou4/GoDirBrowser/handlers"
)

func CreateRoutes() {
	Router.GET("/", handlers.GetDirectory)

	Router.GET("/:path", handlers.GetFile)
	Router.GET("/:path/*subpath", handlers.GetDirectory)

	Router.POST("/upload/*path", handlers.Upload)
}
