package main

import (
	"flag"
	"strconv"

	"github.com/gin-gonic/gin"

	. "github.com/bilou4/GoDirBrowser/constants"
	"github.com/bilou4/GoDirBrowser/routes"
)

func main() {

	var port int = 8080
	var password string
	var ssl bool

	flag.StringVar(&RootDirectory, "directory", InitVar(), "Serve from another directory")
	flag.IntVar(&port, "port", 8080, "Serve from another port than 8080")
	flag.StringVar(&password, "password", "1234", "Password protect the page")
	flag.BoolVar(&ssl, "ssl", false, "Use an SSL connection")

	flag.Parse()

	// Set Gin to production mode
	gin.SetMode(gin.DebugMode)

	// Set the Router as the default one provided by Gin
	Router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded from the disk again.
	Router.LoadHTMLGlob("templates/*")

	// instantiate routes
	routes.CreateRoutes()

	// Start serving the application
	Router.Run(":" + strconv.Itoa(port))

}
