package database

type Database interface {
	AddVideo(video *Video) error
	GetVideo(id string) (Video, error)
	GetVideos() ([]Video, error)
}
