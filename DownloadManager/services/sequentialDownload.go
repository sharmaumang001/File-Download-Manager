package services

import (
	m "DWM/model"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

var Direc []m.Directory

func SeqDownload(url string, ID uuid.UUID) *m.ErrorType {
	urlInfo := strings.Split(url, "/")
	fileName := strings.Split(urlInfo[len(urlInfo)-1], ".")
	extension := fileName[1]

	switch extension {
	case "jpg", "png", "gif":
		response, err := http.Get(url)
		if err != nil {
			return &m.ErrorType{
				ResponseMessage:    "Cannot fetch url",
				ResponseStatusCode: 403,
			}
		}
		defer response.Body.Close()

		file, err := os.Create(filepath.Join(ID.String(), filepath.Base(fileName[0]+"."+extension)))
		if err != nil {
			return &m.ErrorType{
				ResponseMessage:    "Unable to create file",
				ResponseStatusCode: 1003,
			}
		}
		defer file.Close()
		_, err = io.Copy(file, response.Body)
		if err != nil {
			return &m.ErrorType{
				ResponseMessage:    "Unable to save file",
				ResponseStatusCode: 1003,
			}
		}
		var directory m.Directory
		directory.FileName = fileName[0] + "." + extension
		directory.FilePath = "/Users/vikash/Desktop/DownloadManager/app" + "/" + directory.FileName
		Direc = append(Direc, directory)
		return &m.ErrorType{
			ResponseMessage:    "File has been successfully downloaded",
			ResponseStatusCode: 200,
		}
	default:
		return &m.ErrorType{
			ResponseMessage:    "you tried to download wrong extension type file try .jpg  .png  or .gif",
			ResponseStatusCode: 404,
		}
	}

}
