package model

type Err struct {
	ErrMsg  string `json:"err_msg"`
	ErrCode int    `json:"err_code"`
	ErrStr  string `json:"err_str"`
}

type Directory struct {
	FileName string `json:"filename"`
	FilePath string `json:"filepath"`
}
