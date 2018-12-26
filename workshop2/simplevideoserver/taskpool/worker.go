package taskpool

import (
	"github.com/aniskyworker/go-course/workshop2/simplevideoserver/database"
	"github.com/aniskyworker/go-course/workshop2/simplevideoserver/filestorage"
	"github.com/aniskyworker/go-course/workshop2/simplevideoserver/model"
	"github.com/aniskyworker/go-course/workshop2/simplevideoserver/processor"
	log "github.com/sirupsen/logrus"
	"path/filepath"
)

const thumbnailName = "screen.jpg"

func logIfErr(err error) {
	if err != nil {
		log.WithError(err)
	}
}

// Worker create worker for taskpool
func Worker(tasksChan <-chan *Task, db database.Database, videoProcessor processor.VideoProcessor, name int) {
	log.Printf("start worker %v\n", name)
	for task := range tasksChan {
		thumbnailURL := filepath.Join(filestorage.URLContentRoot, task.video.ID, thumbnailName)
		thumbnailFullPath := filepath.Join(filestorage.DirPath, thumbnailURL)
		videoPath := filepath.Join(filestorage.DirPath, task.video.URL)
		err := videoProcessor.CreateVideoThumbnail(videoPath, thumbnailFullPath, 0)
		if err != nil {
			log.WithError(err)
			err = db.UpdateVideoStatus(task.video.ID, model.Error)
			logIfErr(err)
			continue
		}

		dur, err := videoProcessor.GetVideoDuration(videoPath)
		if err != nil {
			log.WithError(err)
			err = db.UpdateVideoStatus(task.video.ID, model.Error)
			logIfErr(err)
			continue
		}

		err = db.UpdateVideo(task.video.ID, thumbnailURL, int(dur))
		if err != nil {
			log.WithError(err)
			err = db.UpdateVideoStatus(task.video.ID, model.Error)
			logIfErr(err)
			continue
		}

		err = db.UpdateVideoStatus(task.video.ID, model.Ready)
		log.WithError(err)
	}
}
