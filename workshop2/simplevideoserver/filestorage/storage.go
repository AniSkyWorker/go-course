package filestorage

import (
	"io"
	"os"
	"path/filepath"
)

// ContentStorage interface of file storage
type ContentStorage interface {
	CreateVideoFile(id string, filename string, content io.Reader) (string, error)
}

const DirPath = "workshop2/simplevideoserver"
const URLContentRoot = "/content"

// VideoStorage video realization of ContentStorage interface
type VideoStorage struct {
}

// CreateVideoFile create video by path that was formed from id
func (vs *VideoStorage) CreateVideoFile(id string, name string, reader io.Reader) (string, error) {
	uniqueFilePath := filepath.Join(URLContentRoot, id)
	file, err := createFile(filepath.Join(DirPath, uniqueFilePath), name)
	if err != nil {
		return "", err
	}

	defer file.Close()

	_, err = io.Copy(file, reader)
	if err != nil {
		return "", err
	}

	return filepath.Join(uniqueFilePath, name), nil
}

func createFile(dirPath string, fileName string) (*os.File, error) {
	if err := os.Mkdir(dirPath, os.ModeDir); err != nil {
		return nil, err
	}
	filePath := filepath.Join(dirPath, fileName)
	return os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
}
