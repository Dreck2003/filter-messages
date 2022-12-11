package helpers

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func Filter[T any](items *[]T, cb func(*T) bool) []T {
	var filteredItems []T
	for i := 0; i < len(*items); i++ {
		if cb(&(*items)[i]) {
			filteredItems = append(filteredItems, (*items)[i])
		}
	}
	return filteredItems
}

func ReadEnvData() (map[string]string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, errors.New("📁 error reading path to current folder")
	}
	path := filepath.Join(pwd, ".env")
	_, notExistError := os.Stat(path)
	if notExistError != nil {
		return nil, errors.New("📁 file .env not find in current directory")
	}

	content, readError := os.ReadFile(path)
	if readError != nil {
		return nil, errors.New("📁 could not read the .env file")
	}
	text_content := strings.Split(string(content), "\n")
	text_content = Filter(&text_content, func(line *string) bool {
		return len(*line) > 0 && strings.Contains(*line, "=")
	})

	content_env := make(map[string]string)

	for i := 0; i < len(text_content); i++ {
		line := text_content[i]
		data := strings.Split(line, "=")
		if len(data) >= 2 {
			content_env[data[0]] = data[1][:len(data[1])-1]
		} else {
			content_env[data[0]] = ""
		}
	}

	return content_env, nil
}
