package controller

import (
	"DWM/model"
	"DWM/services"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	uuid "github.com/google/uuid"
)

type DownloadRequest struct {
	Urls         []string `json:"urls"`
	DownloadType string   `json:"download_type"`
}

var Var DownloadRequest

//creating a ID of string type that we will use for our folder naming
var ID uuid.UUID

func DownloadFiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ID = uuid.New()

	err := os.Mkdir(ID.String(), 0750)
	if err != nil {
		err1 := model.ErrorType{
			ResponseMessage:    "Unable to create file",
			ResponseStatusCode: 404,
		}
		json.NewEncoder(w).Encode(err1.ResponseMessage)
		return
	}

	_ = json.NewDecoder(r.Body).Decode(&Var)
	// fmt.Println("Y is", y)

	urls := Var.Urls
	fmt.Println("hii i m calling", urls)
	DownloadRequestType := Var.DownloadType
	fmt.Println("DownloadType is", DownloadRequestType)
	switch DownloadRequestType {
	case "sequential":
		for i := range urls {
			Url := urls[i]
			err := services.SeqDownload(Url, ID)
			json.NewEncoder(w).Encode(err)
		}
	case "parallel":
		for i := range urls {
			Url := urls[i]
			channel := make(chan model.ErrorType)
			go services.ParaDownload(Url, channel, ID)
			err := <-channel
			json.NewEncoder(w).Encode(err)
		}
	default:
		err := model.ErrorType{
			ResponseMessage:    "Not a valid download type download type can be either parallel or sequential",
			ResponseStatusCode: 404,
		}

		fmt.Println("Not a valid download type download type can be either parallel or sequential", err)

	}
	// json.NewEncoder(w).Encode(ID.String())

}
