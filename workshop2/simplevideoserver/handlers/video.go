package handlers

import (
	"encoding/json"
	"github.com/aniskyworker/go-course/workshop2/simplevideoserver/database"
	"github.com/aniskyworker/go-course/workshop2/simplevideoserver/model"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type videoContent struct {
	ID        string       `json:"id"`
	Name      string       `json:"name"`
	Duration  int          `json:"duration"`
	Thumbnail string       `json:"thumbnail"`
	URL       string       `json:"url"`
	Status    model.Status `json:"status"`
}

func createVideoContent(v model.Video) videoContent {
	return videoContent{v.ID, v.Name, v.Duration, v.Thumbnail, v.URL, v.Status}
}

func getVideo(db database.Database, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]

	video, err := db.GetVideo(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	videoContent := createVideoContent(video)
	b, err := json.Marshal(videoContent)
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
