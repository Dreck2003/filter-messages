package types

type GenericResponse struct {
	Error   *string      `json:"error"`
	Content *interface{} `json:"content"`
}
