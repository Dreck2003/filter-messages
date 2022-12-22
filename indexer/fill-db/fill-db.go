package filldb

import (
	"encoding/json"
	"fmt"

	"github.com/Dreck2003/indexer/config"
	"github.com/Dreck2003/indexer/helpers"
)

const ZINC_URL_BASE = "ZINC_URL"
const ZINC_URL_TO_POST = "/api/_bulkv2"

func PostData(json_data []byte) {
	dataEnv := config.InfoEnvData
	url := string(dataEnv[ZINC_URL_BASE])
	url += ZINC_URL_TO_POST
	req, err := helpers.CreateRequest("POST").AddUrl(url).Build(json_data)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Request.Header.Add("Content-Type", "application/JSON")
	req.Request.SetBasicAuth("admin", "Complexpass#123")
	resp, err := req.Send()
	if err != nil {
		fmt.Println(err)
		return
	}

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
	resp.Body.Close()
}
