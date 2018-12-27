package model

type Video struct {
	ID        string
	Name      string
	Duration  int
	Thumbnail string
	URL       string
	Status    Status
}

// Status video status of video in database
type Status int

const (
	Created    Status = 1
	Processing Status = 2
	Ready      Status = 3
	Deleted    Status = 4
	Error      Status = 5
)
