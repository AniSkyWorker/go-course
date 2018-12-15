package handlers

import (
	"github.com/aniskyworker/go-course/workshop2/simplevideoserver/database"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const dirPath = "workshop2/simplevideoserver"
const urlRoot = "/content"

func uploadVideo(db *database.Database, w http.ResponseWriter, r *http.Request) {
	fileReader, header, err := r.FormFile("file[]")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	contentType := header.Header.Get("Content-Type")
	if contentType != "video/mp4" {
		http.Error(w, err.Error(), http.StatusNoContent)
		return
	}

	videoId := uuid.New()
	fileName := header.Filename
	uniqueFilePath := filepath.Join(urlRoot, videoId.String())
	file, err := createFile(filepath.Join(dirPath, uniqueFilePath), fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer file.Close()
	_, err = io.Copy(file, fileReader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = db.AddVideo(&database.Video{videoId.String(), fileName, 0, "",
		filepath.Join(uniqueFilePath, fileName)})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func createFile(dirPath string, fileName string) (*os.File, error) {
	if err := os.Mkdir(dirPath, os.ModeDir); err != nil {
		return nil, err
	}
	filePath := filepath.Join(dirPath, fileName)
	return os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
}
