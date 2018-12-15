package handlers

import (
	"encoding/json"
	"github.com/aniskyworker/go-course/workshop2/simplevideoserver/database"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func getVideoList(db database.Database, w http.ResponseWriter, _ *http.Request) {
	videos, err := db.GetVideos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(videos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if _, err = io.WriteString(w, string(b)); err != nil {
		log.WithField("err", err).Error("write response error")
	}
}
