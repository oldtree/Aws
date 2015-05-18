package Service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type BrowseService struct {
	List string `json:"FileInfoList"`
}

func (u *BrowseService) ListLogFile(dir string) error {
	var lw = new(LogFileInfo)
	lw.Name = dir
	body, err := lw.Read(dir)
	if err != nil {
		return err
	} else {
		u.List = body
	}
	return nil
}

type LogFileInfo struct {
	IsDir   bool   `json:"IsDir"`
	Size    int64  `json:"Size"`
	Name    string `json:"Name"`
	ModTime string `json:"ModTime"`
}

func (l *LogFileInfo) Read(name string) (string, error) {
	f, err := os.Stat(name)
	var info []LogFileInfo
	if err != nil {
		return "", err
	}
	if f.IsDir() {
		l.IsDir = true
		l.Size = 0
		info = l.ReadDir(name)

		body, err := json.Marshal(&info)
		if err != nil {
			return "", err
		} else {
			return string(body), err
		}
	} else {
		body, err := l.ReadFile(name)
		if err != nil {
			return "", err
		} else {
			return body, nil
		}
	}
}

func (l *LogFileInfo) Format() string {
	b, err := json.Marshal(l)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func (l *LogFileInfo) ReadDir(dir string) []LogFileInfo {
	var filelist []LogFileInfo
	if dir == "" {
		return filelist
	}
	fileinfo, err := ioutil.ReadDir(dir)
	if err != nil {
		return filelist
	}
	for _, value := range fileinfo {
		if value.IsDir() {
			var l = LogFileInfo{
				IsDir:   value.IsDir(),
				Size:    value.Size(),
				Name:    value.Name(),
				ModTime: value.ModTime().String(),
			}
			filelist = append(filelist, l)
		} else {
			var l = LogFileInfo{
				IsDir:   value.IsDir(),
				Size:    value.Size(),
				Name:    value.Name(),
				ModTime: value.ModTime().String(),
			}
			filelist = append(filelist, l)
		}
	}
	return filelist
}

func (l *LogFileInfo) ReadFile(filename string) (string, error) {
	if filename == "" {
		return "", errors.New("filename is not useable")
	}
	f, err := os.Stat(filename)
	if err != nil {
		return "", errors.New("filename is a dir path")
	}
	if f.IsDir() {
		return "", errors.New("filename is a dir path")
	}

	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
