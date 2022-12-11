package config

import (
	"fmt"

	"github.com/Dreck2003/indexer/helpers"
)

var InfoEnvData map[string]string = infoEnvData()

func infoEnvData() map[string]string {
	data, err := helpers.ReadEnvData()
	if err != nil {
		fmt.Println("could not read .env file")
		return make(map[string]string)
	}
	return data
}
