package handlers

import (
	"fmt"
	"net/http"
)

func getVideoList(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, `[{
		"id": "d290f1ee-6c54-4b01-90e6-d701748f0851",
			"name": "Black Retrospective Woman",
			"duration": 15,
			"thumbnail": "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/screen.jpg"
		}]`)
}
