package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Connector struct {
	db *sql.DB
}

func (db *Connector) GetVideos() ([]Video, error) {
	var videos []Video
	rows, err := db.db.Query(`SELECT video_key, title, duration, url, thumbnail_url FROM video`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var video Video
		err := rows.Scan(&video.Id, &video.Name, &video.Duration, &video.Url, &video.Thumbnail)
		if err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}
	return videos, nil
}

func (db *Connector) AddVideo(video *Video) error {
	q := `INSERT INTO video SET video_key = ?, title = ?, duration = ?, url = ?, thumbnail_url = ?`
	_, err := db.db.Exec(q, video.Id, video.Name, video.Duration, video.Url, video.Thumbnail)
	return err
}

func (db *Connector) GetVideo(id string) (Video, error) {
	var video Video
	err := db.db.QueryRow("SELECT video_key, title, duration, url, thumbnail_url FROM video WHERE video_key = ?", id).Scan(
		&video.Id, &video.Name, &video.Duration, &video.Url, &video.Thumbnail)
	return video, err
}

func (db *Connector) Connect() {
	conn, err := sql.Open("mysql", "root:video1234@/videoservice")
	if err != nil {
		log.Fatal(err)
	}
	db.db = conn

	if err := db.db.Ping(); err != nil {
		log.Fatal(err)
	}

}

func (db *Connector) Close() {
	db.Close()
}
