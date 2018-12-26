package taskpool

import (
	"github.com/aniskyworker/go-course/workshop2/simplevideoserver/database"
	"github.com/aniskyworker/go-course/workshop2/simplevideoserver/model"
	"github.com/aniskyworker/go-course/workshop4/videoprocessservice/processor"
	log "github.com/sirupsen/logrus"
	"path/filepath"
)

const thumbnailName = "screen.jpg"

func Worker(tasksChan <-chan *Task, db database.Database, videoProcessor processor.VideoProcessor, thumbnailsDir string, name int) {
	log.Printf("start worker %v\n", name)
	for task := range tasksChan {
		err := db.UpdateVideoStatus(task.video.id, model.Processing)
		if err != nil {
			log.WithError(err)
			continue
		}

		thumbnailPath := filepath.Join(thumbnailsDir, task.video.id, thumbnailName)
		contentPath := "workshop2/simplevideoserver"
		err = videoProcessor.CreateVideoThumbnail(filepath.Join(contentPath, task.video.path), thumbnailPath, 0)
		if err != nil {
			log.WithError(err)
			err = db.UpdateVideoStatus(task.video.id, model.Error)
			if err != nil {
				log.WithError(err)
			}
			continue
		}

		dur, err := videoProcessor.GetVideoDuration(filepath.Join(contentPath, task.video.path))
		if err != nil {
			log.WithError(err)
			err = db.UpdateVideoStatus(task.video.id, model.Error)
			if err != nil {
				log.WithError(err)
			}
			continue
		}

		err = db.UpdateVideo(task.video.id, thumbnailPath, int(dur))
		if err != nil {
			log.WithError(err)
			err = db.UpdateVideoStatus(task.video.id, model.Error)
			if err != nil {
				log.WithError(err)
			}
			continue
		}

		err = db.UpdateVideoStatus(task.video.id, model.Ready)
		log.WithError(err)
	}
}
