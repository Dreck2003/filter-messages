package reader

import (
	"encoding/json"
	"errors"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	filldb "github.com/Dreck2003/indexer/fill-db"
	"github.com/Dreck2003/indexer/helpers"
)

const FinalHeaderEmail string = "X-FileName"
const FinalHeaderLine uint = 8
const INDEX_NAME string = "emails"
const RECORDS string = "records"
const COUNT_PORTION uint = 300

type EmailSection struct {
	num         uint
	typeSection string
}

var SECTIONS_OF_EMAIL = []EmailSection{
	{num: 0, typeSection: "emailId"},
	{num: 1, typeSection: "date"},
	{num: 2, typeSection: "from"},
	{num: 3, typeSection: "to"},
	{num: 4, typeSection: "subject"},
	{num: 8, typeSection: "name"},
}

func GetDataAndFillDB(src string) {
	_, err := os.Stat(src)
	if os.IsNotExist(err) {
		log.Fatal("🧨 The source file not exist 🧨")
	}
	handleError(err)
	readFolder(src)
}

func readFolder(src string) {
	var count uint = 0
	emailsState := []map[string]any{}
	mut := new(sync.Mutex)
	threadPoolToRead := helpers.NewThreadPool(150) // Pool to read files
	threadPoolToSend := helpers.NewThreadPool(3)
	filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		mut.Lock()
		if count >= COUNT_PORTION {
			threadPoolToSend.Execute(func(p ...any) {
				emails := p[0].([]map[string]any)
				structure_email := make(map[string]interface{})
				structure_email["index"] = INDEX_NAME
				structure_email["records"] = emails
				parse, err := parseToJson(structure_email)
				if err != nil {
					return
				}
				filldb.PostData(parse)
			}, emailsState[0:])
			emailsState = make([]map[string]any, 0)
			count = 0
		}
		mut.Unlock()
		threadPoolToRead.Execute(func(p ...any) {
			path := p[0].(string)
			info, errorPath := os.Stat(path)
			if errorPath != nil {
				return
			}
			if !info.IsDir() {
				content, errReadFile := readFile(path)
				if errReadFile != nil {
					return
				}
				string_content, parseError := readContent(content)
				if parseError == nil {
					mut.Lock()
					emailsState = append(emailsState, string_content)
					count = count + 1
					mut.Unlock()
				}
			}
		}, path)
		return nil
	})
	//Wait for two threadSPool:
	threadPoolToRead.Wait()
	threadPoolToSend.Wait()
}

func readFile(path string) (string, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func readContent(content string) (map[string]interface{}, error) {
	lines := strings.Split(content, "\r\n")
	var headerContent []string
	i := 0
	for i < len(lines) {
		if i >= int(FinalHeaderLine) && strings.Contains(lines[i], FinalHeaderEmail) {
			i++
			break
		}
		headerContent = append(headerContent, lines[i])
		i++
	}

	body := strings.Trim(strings.Join(lines[i:], ""), " ")
	return dataToString(headerContent, body)
}

func dataToString(headerContent []string, body string) (map[string]interface{}, error) {
	if len(headerContent) <= 4 {
		return nil, errors.New("the header content over capacity of 5")
	}
	content_json := make(map[string]interface{})
	for i := 0; i < len(SECTIONS_OF_EMAIL); i++ {
		current_section := &SECTIONS_OF_EMAIL[i]
		section_email := strings.Split(headerContent[current_section.num], ":")
		if len(section_email) < 2 {
			content_json[current_section.typeSection] = ""
		} else {
			content_json[current_section.typeSection] = strings.Trim(section_email[1], " ")
		}
	}

	content_json["content"] = body
	return content_json, nil
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func parseToJson(content map[string]interface{}) ([]byte, error) {
	text, err := json.Marshal(content)
	if err != nil {
		return nil, errors.New("cannot posible parse to Json")
	}
	return text, nil
}
