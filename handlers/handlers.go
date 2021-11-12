package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"

	. "github.com/bilou4/GoDirBrowser/constants"
	"github.com/bilou4/GoDirBrowser/utils"
)

// GetFile gets file thanks to the filename in the URL
func GetFile(c *gin.Context) {
	path := c.Param("path")
	subpath := strings.TrimSuffix(c.Param("subpath"), "/")
	if subpath != "" {
		path += subpath
	}

	if file, err := os.Stat(RootDirectory + path); err == nil {
		// path/to/whatever exists
		params := c.Request.URL.Query()
		contentDisposition := "attachment"
		if _, ok := params["view"]; ok {
			contentDisposition = "inline"
		}
		contentType, _ := mimetype.DetectFile(RootDirectory + path)
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("%s; filename=%s", contentDisposition, file.Name()))
		c.Writer.Header().Add("Content-Type", contentType.String())

		c.File(RootDirectory + path)
	} else if errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does *not* exist
		c.HTML(http.StatusNotFound,
			"error.html",
			gin.H{
				"title":   "404 Page",
				"message": "The requested file does not exist.",
			})
	} else {
		// Schrodinger: file may or may not exist. See err for details.
		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
		c.HTML(http.StatusNotFound,
			"error.html",
			gin.H{
				"title":   "404 Page",
				"message": "Error with your request.",
			})
	}
}

// GetDirectory gets directory thanks to the path + subpath in the URL
func GetDirectory(c *gin.Context) {
	path := c.Param("path")
	subpath := strings.TrimSuffix(c.Param("subpath"), "/")

	file, err := os.Stat(RootDirectory + path + subpath)
	if err != nil {
		log.Println(err)
	}
	if file.IsDir() {
		fileList, _ := utils.GetFileList(RootDirectory + path + subpath)
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"title":            "Home Page",
				"file_list":        fileList,
				"parent_directory": path + subpath,
			},
		)
	} else {
		GetFile(c)
	}
}
