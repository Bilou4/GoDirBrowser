package routes

import (
	. "github.com/bilou4/GoDirBrowser/constants"
	"github.com/bilou4/GoDirBrowser/handlers"
	"github.com/gin-gonic/gin"
)

func CreateRoutes(password string) {
	if password != "" {
		authorized := Router.Group("/", gin.BasicAuth(gin.Accounts{
			password: password,
		}))
		authorized.GET("/", handlers.GetDirectory)
		authorized.GET("/:path", handlers.GetFile)
		authorized.GET("/:path/*subpath", handlers.GetDirectory)

		authorized.POST("/upload/*path", handlers.Upload)
	} else {
		Router.GET("/", handlers.GetDirectory)
		Router.GET("/:path", handlers.GetFile)
		Router.GET("/:path/*subpath", handlers.GetDirectory)

		Router.POST("/upload/*path", handlers.Upload)
	}
}
