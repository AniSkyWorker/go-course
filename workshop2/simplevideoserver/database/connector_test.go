package database

import (
	"github.com/aniskyworker/go-course/workshop2/simplevideoserver/model"
	"testing"
)

var videos = []model.Video{
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
}

func TestConnector(t *testing.T) {
	var conn Connector
	conn.Connect("")
	defer conn.Close()

	dbName := "testdb"
	oldDb := "videoservice"
	_, err := conn.db.Exec("CREATE DATABASE " + dbName)

	defer func() {
		_, err = conn.db.Exec("DROP DATABASE " + dbName)
		if err != nil {
			t.Error(err)
			return
		}
	}()

	if err != nil {
		t.Error(err)
		return
	}

	_, err = conn.db.Exec("USE " + dbName)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = conn.db.Exec("CREATE TABLE " + dbName + ".video " + "LIKE " + oldDb + ".video")
	if err != nil {
		t.Error(err)
		return
	}

	for _, video := range videos {
		err = conn.AddVideo(&video)
		if err != nil {
			t.Error(err)
			return
		}

		dbVideo, err := conn.GetVideo(video.ID)
		if err != nil {
			t.Error(err)
			return
		}

		if dbVideo != video {
			t.Errorf("Video from db is wrong. Have: %v, want: %v.", dbVideo, video)
			return
		}
	}

	_, err = conn.GetVideos()
	if err != nil {
		t.Error(err)
		return
	}
}
