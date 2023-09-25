package support

import (
	"fmt"
	"hidden-record-assistant/backend/service/lcu"
	"net/http"
	"os"
	"strings"
)

type FileLoader struct {
	http.Handler
}

func NewFileLoader() *FileLoader {
	return &FileLoader{}
}

func (h *FileLoader) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if strings.HasPrefix(req.URL.Path, "/LCUAPI") {
		url := strings.Replace(req.URL.Path, "/LCUAPI", "", 1)
		bFile, _ := lcu.DefaultApp.Get(url)
		res.Write(bFile)
	} else if strings.HasPrefix(req.URL.Path, "/LOCAL") {
		requestedFilename := strings.Replace(req.URL.Path, "/LOCAL", "./frontend/src", 1)
		fileData, err := os.ReadFile(requestedFilename)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte(fmt.Sprintf("Could not load file %s", requestedFilename)))
		}
		res.Write(fileData)
	}
}
