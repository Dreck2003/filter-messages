package reader

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const FinalHeaderEmail string = "X-FileName"
const FinalHeaderLine uint = 8
const INDEX_NAME string = "emails"
const RECORDS string = "records"

type EmailSection struct {
	num         uint
	typeSection string
}

var SECTIONS_OF_EMAIL = []EmailSection{
	{num: 0, typeSection: "emailId"},
	{num: 2, typeSection: "from"},
	{num: 3, typeSection: "to"},
	{num: 4, typeSection: "subject"},
}

func GetData(src string) {
	_, err := os.Stat(src)
	if os.IsNotExist(err) {
		log.Fatal("The source file not exist")
	}
	handleError(err)
	readFolder(src)
}

func readFolder(src string) {
	count := 0
	emails := []string{}
	filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if count >= 100 {
			structure_email := make(map[string]interface{})
			structure_email["index"] = INDEX_NAME
			structure_email["records"] = emails

			parse, err := parseToJson(structure_email)
			if err != nil {
				return nil
			}
			fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++")
			fmt.Println(string(parse))
			fmt.Println()
			emails = make([]string, 0)
			// count = 0
			time.Sleep(2 * time.Second)
		}
		if count > 500 {
			return errors.New("jAJAJAAJ")
		}

		info, errorPath := os.Stat(path)
		if errorPath != nil {
			return nil
		}
		if !info.IsDir() {
			content, errReadFile := readFile(path)
			if errReadFile != nil {
				return nil
			}
			string_content, parseError := readContent(content)
			if parseError == nil {
				emails = append(emails, string_content)
				count++
			}
		}
		return nil

	})

}

func readFile(path string) (string, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	data := string(body)
	return data, nil
}

func readContent(content string) (string, error) {
	lines := strings.Split(content, "\n")
	var headerContent []string
	i := 0
	for i < len(lines) {
		if i >= int(FinalHeaderLine) && strings.Contains(lines[i], FinalHeaderEmail) {
			break
		}
		headerContent = append(headerContent, lines[i])
		i++
	}

	body := ""
	for i < len(lines) {
		body += lines[i]
		i++
	}
	return dataToString(headerContent, body)
}

func dataToString(headerContent []string, body string) (string, error) {
	// fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++")
	if len(headerContent) <= 4 {
		return "", errors.New("the header content over capacity of 5")
	}
	content_json := make(map[string]interface{})
	for i := 0; i < len(SECTIONS_OF_EMAIL); i++ {
		current_section := &SECTIONS_OF_EMAIL[i]
		section_email := strings.Split(headerContent[current_section.num], ":")
		if len(section_email) < 2 {
			content_json[current_section.typeSection] = ""
		} else {
			content_json[current_section.typeSection] = section_email[1]
		}
	}

	content_json["content"] = body
	return parseToJson(content_json)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func parseToJson(content map[string]interface{}) (string, error) {
	text, err := json.Marshal(content)
	if err != nil {
		return "", errors.New("cannot posible parse to Json")
	}
	return string(text), nil
}
