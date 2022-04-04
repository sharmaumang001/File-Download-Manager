package services

import (
	m "dm/model"
	"dm/utils"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

var Direc []m.Directory

func DownloadSeq(url string, ID uuid.UUID) *m.Err {
	urlInfo := strings.Split(url, "/")
	filename := strings.Split(urlInfo[len(urlInfo)-1], ".")
	extension := filename[1]

	if !utils.IsImage(extension) {
		return &m.Err{
			ErrMsg:  "Not an image",
			ErrCode: 1002,
			ErrStr:  "WRONG_FILE",
		}
	}
	response, err := http.Get(url)

	if err != nil {
		return &m.Err{
			ErrMsg:  "Cannot fetch url",
			ErrCode: 1000,
			ErrStr:  "WRONG_URL",
		}
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return &m.Err{
			ErrMsg:  "Url not giving response",
			ErrCode: 1001,
			ErrStr:  "FAULTY_URL",
		}
	}

	file, err := os.Create(filepath.Join(ID.String(), filepath.Base(filename[0]+"."+extension)))

	if err != nil {
		return &m.Err{
			ErrMsg:  "Unable to create file",
			ErrCode: 1003,
			ErrStr:  "CANNOT_CREATE_FILE",
		}
	}

	defer file.Close()

	_, err = io.Copy(file, response.Body)

	if err != nil {
		return &m.Err{
			ErrMsg:  "Unable to save file",
			ErrCode: 1004,
			ErrStr:  "CANNOT_SAVE_FILE",
		}
	}

	var direc m.Directory
	direc.FileName = filename[0] + "." + extension
	direc.FilePath = "/Users/aashutoshkashyap/Desktop/DownloadManager/app/Downloads/" + filename[0] + "/" + direc.FileName

	Direc = append(Direc, direc)

	return &m.Err{
		ErrMsg:  "File has been successfully downloaded",
		ErrCode: 2000,
		ErrStr:  "",
	}
}
