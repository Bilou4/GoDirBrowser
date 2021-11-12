package utils

import (
	"io/fs"
	"io/ioutil"
)

func GetFileList(root string) ([]fs.FileInfo, error) {
	var files []fs.FileInfo

	files, err := ioutil.ReadDir(root)
	if err != nil {
		return nil, err
	}

	return files, err
}
