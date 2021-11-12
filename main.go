package main

import (
	"github.com/gin-gonic/gin"

	. "github.com/bilou4/GoDirBrowser/constants"
	"github.com/bilou4/GoDirBrowser/routes"
)

func main() {
	// Set Gin to production mode
	gin.SetMode(gin.DebugMode)

	// Set the Router as the default one provided by Gin
	Router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded from the disk again.
	Router.LoadHTMLGlob("templates/*")

	// instantiate routes
	routes.CreateRoutes()

	// Start serving the application
	Router.Run()

}
