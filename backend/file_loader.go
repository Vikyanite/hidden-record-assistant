package backend

import (
	"github.com/Vikyanite/lcu-driver"
	"hidden-record-assistant/backend/module/zlog"
	"net/http"
	"strings"
)

const (
	LCUPrefix = "/lol-game-data"
)

type FileLoader struct {
	http.Handler
}

func NewFileLoader() *FileLoader {
	return &FileLoader{}
}

func (h *FileLoader) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if strings.HasPrefix(req.URL.Path, LCUPrefix) {
		bFile, _ := lcuapi.GET(req.URL.Path)
		res.Write(bFile)
	} else {
		zlog.Errorf("Unhandlered URL: %s", req.URL.Path)
	}
}
