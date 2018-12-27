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

func getVideoStatusByID(id string) *http.Response {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("", "/api/v1/video/"+id, nil)
	r = mux.SetURLVars(r, map[string]string{"ID": id})
	getVideoStatus(&simpleDb, w, r)
	return w.Result()
}

func TestVideoStatus(t *testing.T) {

	videoList, _ := simpleDb.GetVideos()
	for _, video := range videoList {
		r := getVideoStatusByID(video.ID)

		if r.StatusCode != http.StatusOK {
			t.Errorf("Status code is wrong. Have: %d, want: %d.", r.StatusCode, http.StatusOK)
		}

		jsonString, err := ioutil.ReadAll(r.Body)
		r.Body.Close()
		if err != nil {
			t.Fatal(err)
		}
		var status model.Status
		if err = json.Unmarshal(jsonString, &status); err != nil {
			t.Errorf("Can't parse json response with error %v", err)
		}

		if status != model.Ready {
			t.Errorf("Status is wrong. Have: %d, want: %d.", status, model.Ready)
		}
	}

	r := getVideoStatusByID("incorrectid")
	if r.StatusCode != http.StatusInternalServerError {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", r.StatusCode, http.StatusInternalServerError)
	}
}
