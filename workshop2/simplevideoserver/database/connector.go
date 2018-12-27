package database

import (
	"database/sql"
	"github.com/aniskyworker/go-course/workshop2/simplevideoserver/model"
	// mysql syntax driver support
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// Database interface for video management
type Database interface {
	AddVideo(video *model.Video) error
	GetVideo(id string) (model.Video, error)
	GetVideos() ([]model.Video, error)
	GetVideoByStatus(status model.Status) (model.Video, error)
	UpdateVideoStatus(id string, status model.Status) error
	UpdateVideo(id string, thumbnailPath string, duration int) error
}

// Connector mySQL connector that realize Database interface
type Connector struct {
	db *sql.DB
}

// GetVideoByStatus get video by database status
func (db *Connector) GetVideoByStatus(status model.Status) (model.Video, error) {
	var video model.Video
	err := db.db.QueryRow("SELECT video_key, title, duration, url, thumbnail_url FROM video WHERE status = ?", status).Scan(
		&video.ID, &video.Name, &video.Duration, &video.URL, &video.Thumbnail)
	return video, err
}

// UpdateVideoStatus update video status by id
func (db *Connector) UpdateVideoStatus(id string, status model.Status) error {
	_, err := db.db.Query("UPDATE video SET status = ? WHERE video_key = ?", status, id)
	return err
}

// UpdateVideo update video info by id
func (db *Connector) UpdateVideo(id string, thumbnailPath string, duration int) error {
	_, err := db.db.Query("UPDATE video SET thumbnail_url = ?, duration = ? WHERE video_key = ?", thumbnailPath, duration, id)
	return err
}

// GetVideos get all videos from database
func (db *Connector) GetVideos() ([]model.Video, error) {
	var videos []model.Video
	rows, err := db.db.Query(`SELECT video_key, title, duration, url, thumbnail_url FROM video`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var video model.Video
		err := rows.Scan(&video.ID, &video.Name, &video.Duration, &video.URL, &video.Thumbnail)
		if err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}
	return videos, nil
}

// AddVideo add video to database
func (db *Connector) AddVideo(video *model.Video) error {
	q := `INSERT INTO video SET video_key = ?, title = ?, duration = ?, url = ?, thumbnail_url = ?`
	_, err := db.db.Exec(q, video.ID, video.Name, video.Duration, video.URL, video.Thumbnail)
	return err
}

// GetVideo get video from database by id
func (db *Connector) GetVideo(id string) (model.Video, error) {
	var video model.Video
	err := db.db.QueryRow("SELECT video_key, title, duration, url, thumbnail_url, status FROM video WHERE video_key = ?", id).Scan(
		&video.ID, &video.Name, &video.Duration, &video.URL, &video.Thumbnail, &video.Status)
	return video, err
}

// Connect open connection to database by name
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

// Close close connection to database
func (db *Connector) Close() {
	if err := db.db.Close(); err != nil {
		log.Fatal(err)
	}
}
