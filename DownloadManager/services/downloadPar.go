package services

import (
	m "dm/model"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func DownloadPar(url string, ch chan m.Err) {
	urlInfo := strings.Split(url, "/")
	filename := strings.Split(urlInfo[len(urlInfo)-1], ".")
	extension := filename[1]

	if extension != "png" {
		ch <- m.Err{
			ErrMsg:  "Not an image",
			ErrCode: 1002,
			ErrStr:  "WRONG_FILE",
		}
	}
	response, err := http.Get(url)

	if err != nil {
		ch <- m.Err{
			ErrMsg:  "Cannot fetch url",
			ErrCode: 1000,
			ErrStr:  "WRONG_URL",
		}
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		ch <- m.Err{
			ErrMsg:  "Url not giving response",
			ErrCode: 1001,
			ErrStr:  "FAULTY_URL",
		}
	}

	err = os.Mkdir(filename[0], 0750)
	if err != nil {
		ch <- m.Err{
			ErrMsg:  "Unable to create file",
			ErrCode: 1003,
			ErrStr:  "CANNOT_CREATE_FILE",
		}
	}
	file, err := os.Create(filepath.Join(filename[0], filepath.Base(filename[0]+"."+extension)))

	if err != nil {
		ch <- m.Err{
			ErrMsg:  "Unable to create file",
			ErrCode: 1003,
			ErrStr:  "CANNOT_CREATE_FILE",
		}
	}

	defer file.Close()

	_, err = io.Copy(file, response.Body)

	if err != nil {
		ch <- m.Err{
			ErrMsg:  "Unable to save file",
			ErrCode: 1003,
			ErrStr:  "CANNOT_SAVE_FILE",
		}
	}

	var direc m.Directory
	direc.FileName = filename[0] + "." + extension
	direc.FilePath = "/Users/aashutoshkashyap/Desktop/DownloadManager/app/Downloads/" + filename[0] + "/" + direc.FileName

	Direc = append(Direc, direc)

	ch <- m.Err{
		ErrMsg:  "File has been successfully downloaded",
		ErrCode: 2000,
		ErrStr:  "",
	}
}
