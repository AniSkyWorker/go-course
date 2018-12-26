package main

import (
	"github.com/aniskyworker/go-course/workshop2/simplevideoserver/database"
	"github.com/aniskyworker/go-course/workshop2/simplevideoserver/processor"
	"github.com/aniskyworker/go-course/workshop2/simplevideoserver/taskpool"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func RunTaskProvider(stopChan chan struct{}, db database.Database) <-chan *taskpool.Task {
	resultChan := make(chan *taskpool.Task)
	stopTaskProviderChan := make(chan struct{})
	taskProvider := taskpool.TaskProvider{db}
	taskProviderChan := taskProvider.ProvideTasks(stopTaskProviderChan)
	onStop := func() {
		stopTaskProviderChan <- struct{}{}
		close(resultChan)
	}
	go func() {
		for {
			select {
			case <-stopChan:
				onStop()
				return
			case task := <-taskProviderChan:
				select {
				case <-stopChan:
					onStop()
					return
				case resultChan <- task:
				}
			}
		}
	}()
	return resultChan
}

func RunWorkerPool(stopChan chan struct{}, db database.Database, vp processor.VideoProcessor) *sync.WaitGroup {
	var wg sync.WaitGroup
	tasksChan := RunTaskProvider(stopChan, db)
	for i := 0; i < 3; i++ {
		go func(i int) {
			wg.Add(1)
			taskpool.Worker(tasksChan, db, vp, i)
			wg.Done()
		}(i)
	}
	return &wg
}

func main() {
	rand.Seed(time.Now().Unix())
	stopChan := make(chan struct{})

	var db database.Connector
	db.Connect("videoservice")
	defer db.Close()
	var videoProcessor processor.FfmpegVideoProcessor
	killChan := getKillSignalChan()
	wg := RunWorkerPool(stopChan, &db, &videoProcessor)

	waitForKillSignal(killChan)
	stopChan <- struct{}{}
	wg.Wait()
}

func getKillSignalChan() chan os.Signal {
	osKillSignalChan := make(chan os.Signal, 1)
	signal.Notify(osKillSignalChan, os.Kill, os.Interrupt, syscall.SIGTERM)
	return osKillSignalChan
}

func waitForKillSignal(killSignalChan chan os.Signal) {
	killSignal := <-killSignalChan
	switch killSignal {
	case os.Interrupt:
		log.Println("got SIGINT...")
	case syscall.SIGTERM:
		log.Println("got SIGTERM...")
	}
}
