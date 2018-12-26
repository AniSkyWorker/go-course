package handlers

import (
	"encoding/json"
	"errors"
	"github.com/aniskyworker/go-course/workshop2/simplevideoserver/model"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var simpleDb = mockDataBase{
	[]model.Video{
		{ID: "d290f1ee-6c54-4b01-90e6-d701748f0851",
			Name:      "Black Retrospective Woman",
			Duration:  15,
			Thumbnail: "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/screen.jpg",
			URL:       "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/index.mp4"},
		{ID: "sldjfl34-dfgj-523k-jk34-5jk3j45klj34",
			Name:      "Dancing man",
			Duration:  112,
			Thumbnail: "/content/sldjfl34-dfgj-523k-jk34-5jk3j45klj34/screen.jpg",
			URL:       "/content/sldjfl34-dfgj-523k-jk34-5jk3j45klj34/index.mp4"},
		{ID: "hjkhhjk3-23j4-j45k-erkj-kj3k4jl2k345",
			Name:      "Vintage car",
			Duration:  42,
			Thumbnail: "/content/hjkhhjk3-23j4-j45k-erkj-kj3k4jl2k345/screen.jpg",
			URL:       "/content/hjkhhjk3-23j4-j45k-erkj-kj3k4jl2k345/index.mp4"},
	},
	nil,
	nil,
	nil,
}

type mockDataBase struct {
	videos       []model.Video
	getVideoErr  error
	getVideosErr error
	addVideoErr  error
}

func (db *mockDataBase) GetVideos() ([]model.Video, error) {
	return db.videos, db.getVideosErr
}

func (db *mockDataBase) AddVideo(video *model.Video) error {
	db.videos = append(db.videos, *video)
	return db.addVideoErr
}

func (db *mockDataBase) GetVideo(id string) (model.Video, error) {
	for _, video := range db.videos {
		if video.ID == id {
			return video, db.getVideoErr
		}
	}
	return model.Video{}, errors.New("video by specified id not found")
}

func (db *mockDataBase) GetVideoByStatus(status model.Status) (model.Video, error) {
	return model.Video{}, nil
}

func (db *mockDataBase) UpdateVideoStatus(id string, status model.Status) error {
	return nil
}

func (db *mockDataBase) UpdateVideo(id string, thumbnailPath string, duration int) error {
	return nil
}

func getVideos() *http.Response {
	w := httptest.NewRecorder()
	getVideoList(&simpleDb, w, nil)
	return w.Result()
}

func TestList(t *testing.T) {
	response := getVideos()
	if response.StatusCode != http.StatusOK {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", response.StatusCode, http.StatusOK)
	}

	jsonString, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	items := make([]videoContent, 10)
	if err = json.Unmarshal(jsonString, &items); err != nil {
		t.Errorf("Can't parse json response with error %v", err)
	}

	simpleDb.getVideosErr = errors.New("can`t get videos")
	response = getVideos()
	if response.StatusCode != http.StatusInternalServerError {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", response.StatusCode, http.StatusInternalServerError)
	}
}
