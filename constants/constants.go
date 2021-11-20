package constants

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	Router        *gin.Engine
	RootDirectory string
)

func InitVar() string {
	rootDirectory, err := os.Getwd() // directory where uploaded files are stored
	if err != nil {
		log.Println(err)
	}
	return rootDirectory
}
