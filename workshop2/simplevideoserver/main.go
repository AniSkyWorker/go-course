package main

import log "github.com/sirupsen/logrus"

import (
	"fmt"
	"github.com/aniskyworker/go-course/workshop2/simplevideoserver/handlers"
	"net/http"
	"os"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	file, err := os.OpenFile("my.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
		defer file.Close()
	}

	serverUrl := ":8000"
	log.WithFields(log.Fields{"url": serverUrl}).Info("starting the server")
	router := handlers.Router()
	fmt.Println(http.ListenAndServe(":8000", router))
}
