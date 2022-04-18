package model

type ErrorType struct {
	ResponseMessage    string `json:"Response_message"`
	ResponseStatusCode int    `json:"Response_code"`
}
type Directory struct {
	FileName string `json:"filename"`
	FilePath string `json:"filepath"`
}
