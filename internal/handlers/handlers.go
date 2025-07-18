package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "../index.html")
}

func UploadHandler(res http.ResponseWriter, req *http.Request) {

	if err := req.ParseMultipartForm(0); err != nil {
		http.Error(res, "parsing error", http.StatusBadRequest)
	}
	if req.Method != http.MethodPost {
		http.Error(res, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	file, v, err := req.FormFile("myFile")
	if err != nil {
		http.Error(res, "error in receiving file", http.StatusBadRequest)
		return
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(res, "file reading error", http.StatusBadRequest)
		return
	}

	convData := service.Convert(string(data))
	f := filepath.Ext(v.Filename)
	fName := time.Now().UTC().Add(3*time.Hour).Format("02-Jan-2006_15-04-05") + f
	out, err := os.Create(fName)
	if err != nil {
		http.Error(res, "file creating error", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	if _, err := out.WriteString(convData); err != nil {
		http.Error(res, "file writing error", http.StatusInternalServerError)
		return
	}
	if _, err = res.Write([]byte(convData)); err != nil {
		log.Printf("response sending failure: %v", err)
		http.Error(res, "response writing error", http.StatusInternalServerError)
		return
	}
}
