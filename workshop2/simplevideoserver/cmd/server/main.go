package main

import (
	"context"
	"github.com/aniskyworker/go-course/workshop2/simplevideoserver/database"
	"github.com/aniskyworker/go-course/workshop2/simplevideoserver/filestorage"
	"github.com/aniskyworker/go-course/workshop2/simplevideoserver/handlers"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	file, err := os.OpenFile("my.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
		defer file.Close()
	}

	var conn database.Connector
	conn.Connect("videoservice")
	defer conn.Close()

	var videoStorage filestorage.VideoStorage

	serverURL := ":8000"
	log.WithFields(log.Fields{"url": serverURL}).Info("starting the server")
	killSignalChan := getKillSignalChan()
	srv := startServer(serverURL, &conn, &videoStorage)

	waitForKillSignal(killSignalChan)
	err = srv.Shutdown(context.Background())
	if err != nil {
		log.Fatal("Failed to shutdown server")
		return
	}
}

func startServer(serverURL string, conn *database.Connector, vs *filestorage.VideoStorage) *http.Server {
	router := handlers.Router(conn, vs)
	srv := &http.Server{Addr: serverURL, Handler: router}
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	return srv
}

func getKillSignalChan() chan os.Signal {
	osKillSignalChan := make(chan os.Signal, 1)
	signal.Notify(osKillSignalChan, os.Interrupt, syscall.SIGTERM)
	return osKillSignalChan
}

func waitForKillSignal(killSignalChan <-chan os.Signal) {
	killSignal := <-killSignalChan
	switch killSignal {
	case os.Interrupt:
		log.Info("got SIGINT...")
	case syscall.SIGTERM:
		log.Info("got SIGTERM...")
	}
}
