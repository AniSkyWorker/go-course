package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type Video struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Duration  int    `json:"duration"`
	Thumbnail string `json:"thumbnail"`
	Url       string `json:"url"`
}

func getVideo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]

	video := getVideoById(id)
	if video == nil {
		log.Error("can`t find video by " + id)
	} else {
		b, err := json.Marshal(video)
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

}
