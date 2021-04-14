package skbn

import (
	"io"
	"os"
	"path/filepath"
)

// GetListOfFilesFromFs gets list of files in path from Fs (recursive)
func GetListOfFilesFromFs(path string) ([]string, error) {
	var files []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return files, nil
}

// DownloadToFs downloads a single file to Fs
func DownloadToFs(path string, writer io.Writer) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	_, err := f.Write(path)
	if err != nil {
		return err
	}

	return nil
}
