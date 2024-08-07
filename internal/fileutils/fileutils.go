package fileutils

import (
	"os"
	"path/filepath"
)

func GetFileContentString(fileName string) (string, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func HasAllowedExtension(fileName string, extensions []string) bool {
	fileExtension := filepath.Ext(fileName)

	for _, extension := range extensions {
		if fileExtension == extension {
			return true
		}
	}

	return false
}

func ListFileNames(directory string) ([]string, error) {
	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	var result []string
	for _, file := range files {
		result = append(result, file.Name())
	}

	return result, nil
}
