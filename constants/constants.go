package constants

import (
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	Router        *gin.Engine
	RootDirectory string
)

func InitVar() string {
	rootDirectory, err := os.Getwd() // directory where uploaded files are stored
	if err != nil {
		log.Fatal(err)
	}
	return rootDirectory
}
