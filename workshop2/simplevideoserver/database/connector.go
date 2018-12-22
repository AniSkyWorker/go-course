package database

import (
	"database/sql"
	"github.com/aniskyworker/go-course/workshop2/simplevideoserver/model"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Database interface {
	AddVideo(video *model.Video) error
	GetVideo(id string) (model.Video, error)
	GetVideos() ([]model.Video, error)
}

type Connector struct {
	db *sql.DB
}

func (db *Connector) GetVideos() ([]model.Video, error) {
	var videos []model.Video
	rows, err := db.db.Query(`SELECT video_key, title, duration, url, thumbnail_url FROM video`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var video model.Video
		err := rows.Scan(&video.Id, &video.Name, &video.Duration, &video.Url, &video.Thumbnail)
		if err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}
	return videos, nil
}

func (db *Connector) AddVideo(video *model.Video) error {
	q := `INSERT INTO video SET video_key = ?, title = ?, duration = ?, url = ?, thumbnail_url = ?`
	_, err := db.db.Exec(q, video.Id, video.Name, video.Duration, video.Url, video.Thumbnail)
	return err
}

func (db *Connector) GetVideo(id string) (model.Video, error) {
	var video model.Video
	err := db.db.QueryRow("SELECT video_key, title, duration, url, thumbnail_url FROM video WHERE video_key = ?", id).Scan(
		&video.Id, &video.Name, &video.Duration, &video.Url, &video.Thumbnail)
	return video, err
}

func (db *Connector) Connect(dbName string) {
	conn, err := sql.Open("mysql", "root:video1234@/"+dbName)
	if err != nil {
		log.Fatal(err)
	}
	db.db = conn

	if err := db.db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func (db *Connector) Close() {
	if err := db.db.Close(); err != nil {
		log.Fatal(err)
	}
}
