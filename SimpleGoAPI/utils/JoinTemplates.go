package utils

import (
	"fmt"
	"path/filepath"
)

//This function gather all of HTML files and serve in a full package
func Include() []string {
	files, err := filepath.Glob("view/*.html")
	if err != nil {
		fmt.Print(err)
	}
	for _, file := range files {
		files = append(files, file)
	}
	return files
}
