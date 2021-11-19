package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	. "github.com/bilou4/GoDirBrowser/constants"
	"github.com/bilou4/GoDirBrowser/handlers"
	"github.com/bilou4/GoDirBrowser/utils"
)

func CreateRoutes() {
	Router.GET("/", func(c *gin.Context) {
		fileList, err := utils.GetFileList(".")
		if err != nil {
			panic("error listing files.")
		}
		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title":             "Home Page",
				"file_list":         fileList,
				"current_directory": RootDirectory,
			},
		)
	})

	Router.GET("/:path", handlers.GetFile)
	Router.GET("/:path/*subpath", handlers.GetDirectory)
}
