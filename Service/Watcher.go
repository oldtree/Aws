package Service

import (
	"encoding/json"
	"time"
)

type ServerStatu struct {
	InNumber       int
	OkOutNumber    int
	ErrorOutNumber int
}

type AppStaticsService struct {
	StartTime string
	StopTime  string
	AppServ   map[string]*ServerStatu
}

func NewAppStaticsService() *AppStaticsService {
	a := new(AppStaticsService)
	a.StartTime = time.Now().String()
	a.AppServ = make(map[string]*ServerStatu)
	a.AppServ["blog"] = &ServerStatu{
		0,
		0,
		0,
	}
	a.AppServ["history"] = &ServerStatu{
		0,
		0,
		0,
	}
	return a
}

func (u *AppStaticsService) Format() string {
	str, err := json.Marshal(u)
	if err != nil {
		return ""
	}
	return string(str)
}

func (u *AppStaticsService) GetAppServiceStatics() (msg string, err error) {
	str, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(str), nil
}
