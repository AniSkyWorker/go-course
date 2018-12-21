package filestorage

import (
	"os"
	"path/filepath"
	"testing"
)

func TestVideoUpload(t *testing.T) {
	var vs VideoStorage
	file, err := os.Open("../content/d290f1ee-6c54-4b01-90e6-d701748f0851/index.mp4")
	if err != nil {
		t.Error(err)
		return
	}
	defer file.Close()

	dir, id, name := "../content/", "some_id", "test_file"
	path, err := vs.CreateVideoFile(id, name, file)
	if err != nil {
		t.Error(err)
		return
	}
	path = filepath.Join("..", path)
	if _, err = os.Stat(path); os.IsNotExist(err) {
		t.Error(err)
		return
	}

	if err = os.RemoveAll(filepath.Join(dir, id)); err != nil {
		t.Error("test file delete failed")
		return
	}
}
