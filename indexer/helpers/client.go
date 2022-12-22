package helpers

import (
	"net/http"
	"time"
)

var NetClient = &http.Client{
	Timeout: time.Second * 30,
	//more features to be included
}
