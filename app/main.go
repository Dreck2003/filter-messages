package main

import (
	"log"
	"net/http"

	"github.com/Dreck2003/api/backend"
)

func main() {
	err := http.ListenAndServe(":3000", backend.CreateRoutes())
	if err != nil {
		log.Fatal(err)
	}
}
