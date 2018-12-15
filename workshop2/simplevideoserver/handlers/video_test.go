package handlers

import (
	context2 "context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVideo(t *testing.T) {

	for _, video := range videoList {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("", "http://example.com/", nil)
		context := context2.WithValue(context2.Background(), "ID", video.Id)
		r = r.WithContext(context)
		getVideo(w, r)
		response := w.Result()

		if response.StatusCode != http.StatusOK {
			t.Errorf("Status code is wrong. Have: %d, want: %d.", response.StatusCode, http.StatusOK)
		}

		jsonString, err := ioutil.ReadAll(response.Body)
		response.Body.Close()
		if err != nil {
			t.Fatal(err)
		}
		var video Video
		if err = json.Unmarshal(jsonString, &video); err != nil {
			t.Errorf("Can't parse json response with error %v", err)
		}
	}
}
