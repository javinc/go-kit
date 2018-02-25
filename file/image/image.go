package image

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
)

// image manipulation

// Resize creates a scaled image with new dimensions
// If either width or height is set to 0, it will be set to an aspect ratio preserving value
func Resize(path string, width, height uint) (newPath string, err error) {
	return
}

// Thumbnail downscales an image preserving its aspect ratio to the maximum dimensions
//  It will return the original image if original sizes are smaller than the provided dimensions.
func Thumbnail(path string, width, height uint) (newPath string, err error) {
	// check for cache

	img, err := decodeImage(path)
	if err != nil {
		return
	}

	_, filename := filepath.Split(path)
	m := resize.Thumbnail(width, height, img, resize.Lanczos3)
	out, newPath, err := Cache(filename, width, height)
	if err != nil {
		return
	}
	defer out.Close()

	// write new image to file
	encodeImage(newPath, out, m)

	return
}

func encodeImage(newPath string, file *os.File, img image.Image) error {
	ext := filepath.Ext(newPath)
	switch strings.ToUpper(ext) {
	case ".JPG":
	case ".JPEG":
		return jpeg.Encode(file, img, nil)
	case ".PNG":
		return png.Encode(file, img)
	case ".GIF":
		return gif.Encode(file, img, nil)
	default:
		return errors.New("unsupported image type " + ext)
	}

	return nil
}

func decodeImage(path string) (image.Image, error) {
	// open file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// decode base on type
	var img image.Image
	ext := filepath.Ext(path)
	switch strings.ToUpper(ext) {
	case ".JPG":
	case ".JPEG":
		img, err = jpeg.Decode(file)
		if err != nil {
			return nil, err
		}
	case ".PNG":
		img, err = png.Decode(file)
		if err != nil {
			return nil, err
		}
	case ".GIF":
		img, err = gif.Decode(file)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("unsupported image type " + ext)
	}

	return img, nil
}
