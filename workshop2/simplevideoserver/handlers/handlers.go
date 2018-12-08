package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/api/v1/list", getVideoList).Methods(http.MethodGet)
	s.HandleFunc("/api/v1/video/d290f1ee-6c54-4b01-90e6-d701748f0851", getVideo).Methods(http.MethodGet)
	return r
}
