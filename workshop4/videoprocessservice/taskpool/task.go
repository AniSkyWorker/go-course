package taskpool

type Video struct {
	id   string
	path string
}

type Task struct {
	video Video
}
