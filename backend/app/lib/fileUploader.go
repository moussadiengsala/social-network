package lib

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// ImageUploader handles image uploads
func FileUploader(r *http.Request, tag string) (string, error) {
	uploadsDir := "./static/uploads/"
	// Check if the uploads directory exists
	_, err := os.Stat(uploadsDir)
	if os.IsNotExist(err) {
		os.MkdirAll(uploadsDir, os.ModePerm)
	}

	file, handler, err := r.FormFile(tag)
	if err != nil {
		if err == http.ErrMissingFile {
			// No file uploaded, return without error
			return "", nil
		}
		// Handle other errors
		return "", err
	}

	supportedExtensions := []string{".jpeg", ".png", ".gif", ".jpg"}
	extensionNotSupported := true
	for _, supportedExt := range supportedExtensions {
		if filepath.Ext(handler.Filename) == supportedExt {
			extensionNotSupported = false
			break
		}
	}
	if extensionNotSupported {
		return "", errors.New("unsupported image extension")
	}

	for _, fileHeaders := range r.MultipartForm.File {
		for _, fileHeader := range fileHeaders {
			// Check the file size.
			if fileHeader.Size > (20 * 1024 * 1024) {
				return "", errors.New("file size exceeds the limit.")
			}
		}
	}

	fileName := time.Now().Format("20060102150405") + filepath.Ext(handler.Filename)
	filepath := filepath.Join(uploadsDir, fileName)

	f, err := os.Create(filepath)
	if err != nil {
		// Handle the error appropriately
		return "", err
	}
	defer f.Close()

	_, copyErr := io.Copy(f, file)
	if copyErr != nil {
		// Handle the error appropriately
		return "", copyErr
	}

	return fileName, nil
}
