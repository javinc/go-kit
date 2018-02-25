package image

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/nfnt/resize"
)

// image manipulation

var (
	// Path upload directory
	Path = "./upload/cache/"
)

// Resize creates a scaled image with new dimensions
// If either width or height is set to 0, it will be set to an aspect ratio preserving value
func Resize(path string, width, height uint) (newPath string, err error) {
	return
}

// Thumbnail downscales an image preserving its aspect ratio to the maximum dimensions
//  It will return the original image if original sizes are smaller than the provided dimensions.
func Thumbnail(path string, width, height uint) (newPath string, err error) {
	newPath = composePath(path, width, height)

	// check for cache

	// make cache path writable
	dir, _ := filepath.Split(newPath)
	os.MkdirAll(dir, 0777)
	if err != nil {
		return
	}

	img, err := decodeImage(path)
	if err != nil {
		return
	}

	m := resize.Thumbnail(width, height, img, resize.Lanczos3)
	out, err := os.Create(newPath)
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
		fallthrough
	case ".JPEG":
		return jpeg.Encode(file, img, nil)
	case ".PNG":
		return png.Encode(file, img)
	case ".GIF":
		return gif.Encode(file, img, nil)
	default:
		return errors.New("unsupported image type " + ext)
	}
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
		fallthrough
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

// outputs /upload/cache/600x400/test.jpg
func composePath(filename string, width, height uint) string {
	strconv.Itoa(int(width))

	// clean upload path
	Path = strings.TrimSuffix(Path, "/")
	dimension := strconv.Itoa(int(width)) + "x" + strconv.Itoa(int(height))

	return strings.Join([]string{
		Path,
		dimension,
		filename,
	}, "/")
}
