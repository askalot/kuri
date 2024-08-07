package fileutils

import (
	"os"
	"path"
	"path/filepath"
)

func CreateFileWithContentString(fileName string, content string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func GetFileContentString(fileName string) (string, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func GetFileNameWithoutExtension(fileName string) string {
	return fileName[:len(fileName)-len(path.Ext(fileName))]
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
