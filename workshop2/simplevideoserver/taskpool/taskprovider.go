package taskpool

import (
	"github.com/aniskyworker/go-course/workshop2/simplevideoserver/database"
	"github.com/aniskyworker/go-course/workshop2/simplevideoserver/model"
	log "github.com/sirupsen/logrus"
	"time"
)

func (tp *TaskProvider) GenerateTask() *Task {
	video, err := tp.Database.GetVideoByStatus(model.Created)
	if err != nil {
		log.WithError(err)
		return nil
	}
	err = tp.Database.UpdateVideoStatus(video.ID, model.Processing)
	if err != nil {
		log.WithError(err)
		return nil
	}

	return &Task{video}
}

type TaskProvider struct {
	Database database.Database
}

func (tp *TaskProvider) ProvideTasks(stopChan chan struct{}) <-chan *Task {
	tasksChan := make(chan *Task)
	go func() {
		for {
			select {
			case <-stopChan:
				close(tasksChan)
				return
			default:
			}
			if task := tp.GenerateTask(); task != nil {
				log.Printf("got the task %v\n", task)
				tasksChan <- task
			} else {
				log.Println("no task for processing, start waiting")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return tasksChan
}
