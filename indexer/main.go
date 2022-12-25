package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/Dreck2003/indexer/commands"
	"github.com/Dreck2003/indexer/reader"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:7000", nil)) // Serve only for profiling :v
	}()
	command := commands.GetArgs()
	reader.GetDataAndFillDB(command)

}
