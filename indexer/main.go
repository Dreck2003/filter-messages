package main

import (
	"github.com/Dreck2003/indexer/commands"
	"github.com/Dreck2003/indexer/reader"
)

func main() {
	command := commands.GetArgs()
	reader.GetData(command)
}
