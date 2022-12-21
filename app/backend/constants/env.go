package constants

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/Dreck2003/api/backend/utils"
)

var Env = GetEnv()

var (
	ZINC_USERNAME = "ZINC_USERNAME"
	ZINC_PASSWORD = "ZINC_PASSWORD"
	ZINC_URL      = "ZINC_URL"
)

func GetEnv() map[string]string {
	mapCount := make(map[string]string)
	env := os.Environ()

	for _, line := range env {
		lineSplit := strings.Split(line, "=")
		if len(lineSplit) >= 2 {
			mapCount[lineSplit[0]] = lineSplit[1]
		}
	}

	envData, err := ReadEnvFile()
	if err == nil {
		for key, val := range envData {
			mapCount[key] = val
		}
	}
	return mapCount
}

func ReadEnvFile() (map[string]string, error) {
	envGroup := make(map[string]string)
	root, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	pathToEnvFile := filepath.Join(root, ".env")
	file, err := os.Open(pathToEnvFile)
	if err != nil {
		fmt.Println("Cannot find .env file")
		return nil, err
	}
	defer file.Close()
	content, err := fs.ReadFile(os.DirFS(root), ".env")
	if err != nil {
		fmt.Println("Cannot read the .env file")
		return nil, err
	}
	dataContent := string(content)
	lines := strings.Split(dataContent, "\n")
	for _, line := range lines {
		substr := utils.FirstIndexSkipping(line, " ")
		if substr == "#" { // So this line is a comment
			continue
		}
		lineSplit := strings.Split(line, "=")
		if len(lineSplit) >= 2 {
			cleanFirst := strings.Trim(strings.Trim(lineSplit[0], " "), "\r")
			cleanSecond := strings.Trim(strings.Trim(lineSplit[1], " "), "\r")

			if cleanSecond[0:1] == `"` {
				cleanSecond = utils.SubstrWithEnd(cleanSecond, `"`)
				cleanSecond = strings.Trim(cleanSecond, `"`)
			} else {
				haveComment := strings.Contains(cleanSecond, "#")
				if haveComment {
					index := strings.Index(cleanSecond, "#")
					if index != -1 {
						cleanSecond = strings.Trim(cleanSecond[0:index], " ")
					}
				}
			}
			envGroup[cleanFirst] = strings.Trim(cleanSecond, `'`)
		}
	}
	return envGroup, nil
}
