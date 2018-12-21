package handlers

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/aniskyworker/go-course/workshop2/simplevideoserver/filestorage"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"testing"
)

func mockFileRequest(path string) *http.Request {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "file[]", path))
	h.Set("Content-Type", "video/mp4")

	_, err := writer.CreatePart(h)
	if err != nil {
		return nil
	}

	err = writer.Close()
	if err != nil {
		return nil
	}

	request := httptest.NewRequest("", "/api/v1/video", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	return request
}

type mockFileStorage struct {
	createVideoError error
}

func (fs *mockFileStorage) CreateVideoFile(id string, name string, reader io.Reader) (string, error) {
	return id, fs.createVideoError
}

func uploadVideoRequest(fs filestorage.ContentStorage, r *http.Request) *http.Response {
	w := httptest.NewRecorder()
	uploadVideo(&simpleDb, fs, w, r)
	return w.Result()
}

func TestVideoUpload(t *testing.T) {
	var mockFS mockFileStorage
	request := httptest.NewRequest("", "/api/v1/video", nil)
	result := uploadVideoRequest(&mockFS, request)
	if result.StatusCode != http.StatusInternalServerError {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", result.StatusCode, http.StatusInternalServerError)
	}

	request = mockFileRequest("exampleVideo")
	result = uploadVideoRequest(&mockFS, request)
	if result.StatusCode != http.StatusOK {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", result.StatusCode, http.StatusOK)
	}

	mockFS.createVideoError = errors.New("can`t create video file")
	request = mockFileRequest("exampleVideo")
	result = uploadVideoRequest(&mockFS, request)
	if result.StatusCode != http.StatusInternalServerError {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", result.StatusCode, http.StatusInternalServerError)
	}
}
