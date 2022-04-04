package controller

import (
	"dm/model"
	"dm/services"
	"encoding/json"
	"net/http"
	"os"

	"github.com/google/uuid"
)

type DownloadRequest struct {
	Urls         []string `json:"urls"`
	DownloadType string   `json:"download_type"`
}

var Err []model.Err

var ID uuid.UUID

func DownloadFileSeq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var dr DownloadRequest
	ID = uuid.New()

	err := os.Mkdir(ID.String(), 0750)
	if err != nil {
		err1 := model.Err{
			ErrMsg:  "Unable to create file",
			ErrCode: 1003,
			ErrStr:  "CANNOT_CREATE_FILE",
		}
		json.NewEncoder(w).Encode(err1.ErrMsg)
		return
	}

	_ = json.NewDecoder(r.Body).Decode(&dr)

	urls := dr.Urls
	downloadType := dr.DownloadType

	if downloadType == "sequential" {
		for i := range urls {
			Url := urls[i]
			err := services.DownloadSeq(Url, ID)
			json.NewEncoder(w).Encode(err.ErrMsg)
		}
	} else if downloadType == "parallel" {
		for i := range urls {
			Url := urls[i]
			ch := make(chan model.Err)
			go services.DownloadPar(Url, ch, ID)
			err := <-ch
			json.NewEncoder(w).Encode(err)
		}
	} else {

		err := model.Err{
			ErrMsg:  "Not a valid download type. Download type can be either 'sequential' or 'parallel'.",
			ErrCode: 1005,
			ErrStr:  "INVALID_DOWNLOAD_TYPE",
		}
		json.NewEncoder(w).Encode(err.ErrMsg)
	}
	json.NewEncoder(w).Encode("Download ID:" + ID.String())
}
