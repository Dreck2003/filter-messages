package filldb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Dreck2003/indexer/config"
)

const ZINC_URL_BASE = "ZINC_URL"
const ZINC_URL_TO_POST = "/api/_bulkv2"

func PostData(json_data []byte) {
	dataEnv := config.InfoEnvData
	url := string(dataEnv[ZINC_URL_BASE])
	url += ZINC_URL_TO_POST
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json_data))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/JSON")
	req.SetBasicAuth("admin", "Complexpass#123")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		var data map[string]interface{}
		fmt.Println()
		if resp.StatusCode != 200 {
			fmt.Println("-------- 🧨 Error sending data to ZincSearch 😥----------------")
			fmt.Println(">> StatusCode", resp.StatusCode)
			json.NewDecoder(resp.Body).Decode(&data)
			fmt.Println(data)
		} else {
			fmt.Println("--------🚀 Sending data correctly ✨ ----------------")
		}
	}
	resp.Body.Close()
}
