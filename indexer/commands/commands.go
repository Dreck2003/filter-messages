package commands

import (
	"log"
	"os"
	"strings"
)

func GetArgs() string {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("The src folder not exist!")
	}
	return strings.Trim(args[1], " ")

}
