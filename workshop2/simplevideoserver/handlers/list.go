package handlers

import (
	"encoding/json"
	"github.com/aniskyworker/go-course/workshop2/simplevideoserver/database"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func getVideoList(db database.Database, w http.ResponseWriter, r *http.Request) {
	videos, err := db.GetVideos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	q := r.URL.Query()
	searchByName := q.Get("searchString")
	skip, err := strconv.Atoi(q.Get("skip"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	limit, err := strconv.Atoi(q.Get("limit"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var videoContentList []videoContent
	for i, video := range videos {
		if strings.Contains(video.Name, searchByName) && i >= skip && i < skip+limit {
			videoContentList = append(videoContentList, createVideoContent(video))
		}
	}

	b, err := json.Marshal(videoContentList)
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
