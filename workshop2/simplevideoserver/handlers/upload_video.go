package handlers

import (
	"github.com/aniskyworker/go-course/workshop2/simplevideoserver/database"
	"github.com/aniskyworker/go-course/workshop2/simplevideoserver/filestorage"
	"github.com/aniskyworker/go-course/workshop2/simplevideoserver/model"
	"github.com/google/uuid"
	"net/http"
)

func uploadVideo(db database.Database, cs filestorage.ContentStorage, w http.ResponseWriter, r *http.Request) {
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

	uniqueFilePath, err := cs.CreateVideoFile(videoId.String(), fileName, fileReader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = db.AddVideo(&model.Video{Id: videoId.String(), Name: fileName, Duration: 0, Thumbnail: "",
		Url: uniqueFilePath})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
