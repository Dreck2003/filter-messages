package main

import (
	"log"
	"sync"

	"net/http"
	_ "net/http/pprof"

	"github.com/Dreck2003/indexer/commands"
	"github.com/Dreck2003/indexer/reader"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:7000", nil)) // Serve only for profiling :v
	}()
	wg := sync.WaitGroup{} // Create a sync.WaitGroup to wait for resolve all goroutines
	command := commands.GetArgs()
	reader.GetDataAndFillDB(command, &wg)
	wg.Wait() // Wait to resolve goroutines
}
