package handlers

import (
	"github.com/aniskyworker/go-course/workshop2/simplevideoserver/database"
	"github.com/aniskyworker/go-course/workshop2/simplevideoserver/filestorage"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func WrapHandlerWithDb(db database.Database, f func(db database.Database, w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(db, w, r)
	}
}

func WrapHandlerWithVideoStorage(db database.Database, vs filestorage.ContentStorage, f func(db database.Database, cs filestorage.ContentStorage, w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return WrapHandlerWithDb(db, func(db database.Database, w http.ResponseWriter, r *http.Request) {
		f(db, vs, w, r)
	})
}

func Router(db database.Database, vs filestorage.ContentStorage) http.Handler {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()

	s.HandleFunc("/list", WrapHandlerWithDb(db, getVideoList)).Methods(http.MethodGet)
	s.HandleFunc("/video/{ID}", WrapHandlerWithDb(db, getVideo)).Methods(http.MethodGet)
	s.HandleFunc("/video", WrapHandlerWithVideoStorage(db, vs, uploadVideo)).Methods(http.MethodPost)

	return logMiddleware(r)
}

func logMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"method":     r.Method,
			"url":        r.URL,
			"remoteAddr": r.RemoteAddr,
			"userAgent":  r.UserAgent(),
		}).Info("got a new request")
		h.ServeHTTP(w, r)
	})
}
