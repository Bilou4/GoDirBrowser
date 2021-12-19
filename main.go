package main

import (
	"embed"
	"flag"
	"html/template"
	"log"
	"os"
	"strconv"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"

	. "github.com/bilou4/GoDirBrowser/constants"
	"github.com/bilou4/GoDirBrowser/routes"
)

//go:embed templates/*
var templatesFS embed.FS

func main() {

	var port int = 8080
	var password string
	var ssl bool

	flag.StringVar(&RootDirectory, "directory", InitVar(), "Serve from another directory")
	flag.IntVar(&port, "port", 8080, "Serve from another port than 8080")
	flag.StringVar(&password, "password", "", "Password protect the page")
	flag.BoolVar(&ssl, "ssl", false, "Use an SSL connection")

	flag.Parse()

	if c := RootDirectory[len(RootDirectory)-1]; string(c) != "/" {
		RootDirectory += "/"
	}

	// Set Gin to production mode
	gin.SetMode(gin.DebugMode)

	// Set the Router as the default one provided by Gin
	Router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded from the disk again.
	render := multitemplate.New()

	list := [...]string{
		"error.html",
		"index.html",
	}

	for _, x := range list {
		data, err := templatesFS.ReadFile("templates/" + x)
		if err != nil {
			log.Fatal(err)
		}

		tmplMessage, err := template.New(x).Parse(string(data))
		if err != nil {
			log.Println(err)
		}
		render.Add(x, tmplMessage)
	}
	Router.HTMLRender = render

	// instantiate routes
	routes.CreateRoutes(password)

	// Start serving the application
	if ssl {
		certificateFileName := "certificate.crt"
		certificateKeyFileName := "certificate_key.key"
		if _, err := os.Stat(certificateFileName); err != nil {
			log.Fatalf("Cannot find certificate file")
		}
		if _, err := os.Stat(certificateKeyFileName); err != nil {
			log.Fatalf("Cannot find certificate key file")
		}
		Router.RunTLS(":"+strconv.Itoa(port), certificateFileName, certificateKeyFileName)
	} else {
		Router.Run(":" + strconv.Itoa(port))
	}

}
