package main

import (
	"fmt"
	"github.com/aniskyworker/go-course/workshop2/simplevideoserver/handlers"
	"net/http"
)

func main() {
	router := handlers.Router()
	fmt.Println(http.ListenAndServe(":8000", router))
}
