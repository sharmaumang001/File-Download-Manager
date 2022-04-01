package controller

import (
	"dm/model"
	"dm/services"
	"encoding/json"
	"net/http"
)

type DownloadRequest struct {
	Urls         []string `json:"urls"`
	DownloadType string   `json:"download_type"`
}

func DownloadFileSeq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var dr DownloadRequest

	_ = json.NewDecoder(r.Body).Decode(&dr)

	urls := dr.Urls
	downloadType := dr.DownloadType

	if downloadType == "sequential" {
		for i := range urls {
			Url := urls[i]
			err := services.DownloadSeq(Url)
			json.NewEncoder(w).Encode(err.ErrMsg)
		}
	} else if downloadType == "parallel" {
		for i := range urls {
			Url := urls[i]
			ch := make(chan model.Err)
			go services.DownloadPar(Url, ch)
			err := <-ch
			json.NewEncoder(w).Encode(err.ErrMsg)
		}
	} else {

		err := model.Err{
			ErrMsg:  "Not a valid download type. Download type can be either 'sequential' or 'parallel'.",
			ErrCode: 1005,
			ErrStr:  "INVALID_DOWNLOAD_TYPE",
		}
		json.NewEncoder(w).Encode(err.ErrMsg)
	}
}
