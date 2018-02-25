package image

import (
	"os"
	"strconv"
	"strings"
)

// image caching

var (
	// Path upload directory
	Path = "./upload/cache/"
)

// Cache by dimension
func Cache(filename string, width, height uint) (file *os.File, newPath string, err error) {
	newPath = composePath(filename, width, height)

	// make upload path writable
	os.Mkdir(newPath, 0777)
	if err != nil {
		return
	}

	out, err := os.Create(newPath + filename)
	if err != nil {
		return
	}
	defer out.Close()

	return
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
