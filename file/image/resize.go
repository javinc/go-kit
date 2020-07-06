package image

import (
	"os"
	"path/filepath"

	"github.com/h2non/bimg"
)

// Resize requires libvips see for installation https://github.com/libvips/libvips
func Resize(path string, width, height uint) (newPath string, err error) {
	newPath = composePath(path, width, height)

	// Check for cache
	if _, exists := os.Stat(newPath); exists == nil {
		return
	}

	// Make cache path writable
	dir, _ := filepath.Split(newPath)
	if err = os.MkdirAll(dir, os.ModePerm); err != nil {
		return
	}

	buffer, err := bimg.Read(path)
	if err != nil {
		return
	}

	newImage, err := bimg.NewImage(buffer).Resize(int(width), int(height))
	if err != nil {
		return
	}

	if err = bimg.Write(newPath, newImage); err != nil {
		return
	}

	return newPath, nil
}
