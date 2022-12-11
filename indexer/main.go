package main

import (
	"sync"

	"github.com/Dreck2003/indexer/commands"
	"github.com/Dreck2003/indexer/reader"
)

func main() {

	wg := sync.WaitGroup{} // Create a sync.WaitGroup to wait for resolve all goroutines
	command := commands.GetArgs()
	reader.GetDataAndFillDB(command, &wg)
	wg.Wait() // Wait to resolve goroutines
}
