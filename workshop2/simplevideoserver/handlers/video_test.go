package handlers

import (
	"encoding/json"
	"github.com/aniskyworker/go-course/workshop2/simplevideoserver/model"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getVideoById(id string) *http.Response {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("", "/api/v1/video/"+id, nil)
	r = mux.SetURLVars(r, map[string]string{"ID": id})
	getVideo(&simpleDb, w, r)
	return w.Result()
}

func TestVideo(t *testing.T) {

	videoList, _ := simpleDb.GetVideos()
	for _, video := range videoList {
		r := getVideoById(video.Id)

		if r.StatusCode != http.StatusOK {
			t.Errorf("Status code is wrong. Have: %d, want: %d.", r.StatusCode, http.StatusOK)
		}

		jsonString, err := ioutil.ReadAll(r.Body)
		r.Body.Close()
		if err != nil {
			t.Fatal(err)
		}
		var video model.Video
		if err = json.Unmarshal(jsonString, &video); err != nil {
			t.Errorf("Can't parse json response with error %v", err)
		}

		r = getVideoById("incorrectid")
		if r.StatusCode != http.StatusInternalServerError {
			t.Errorf("Status code is wrong. Have: %d, want: %d.", r.StatusCode, http.StatusInternalServerError)
		}
	}
}
