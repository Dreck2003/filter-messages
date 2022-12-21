package types

import "fmt"

var BaseErrors = map[string]string{
	"EmptyString": "The provided value is empty",
}

type HttpErrorsType struct{}

var HttpErrors = HttpErrorsType{}

func (e *HttpErrorsType) NotProvided(val string) string {
	return fmt.Sprintf("The value: %s was not provided", val)
}
