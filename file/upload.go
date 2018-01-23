package file

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"

	"github.com/javinc/go-kit/hash"
)

var (
	// Path upload directory
	Path = "./upload/"
	// Size upload limit
	Size = 5242880
	// Mime type allowed
	Mime = "image/"
)

// Upload file
func Upload(file multipart.File, header *multipart.FileHeader, err error) (string, string, error) {
	var name string

	// clean upload path
	Path = strings.TrimSuffix(Path, "/") + "/"
	// make upload path writable
	os.Mkdir(Path, 0777)
	if err != nil {
		return name, "FILE_UPLOAD_ERROR", err
	}

	size, err := getFileSize(file)
	if err != nil {
		return name, "FILE_UPLOAD_SIZE_ERROR", err
	}

	// check size limit
	if Size < size {
		return name, "FILE_UPLOAD_SIZE_LIMIT", errors.New(strconv.Itoa(Size/1000000) + "MB upload size limit reached")
	}

	// check mime type
	mime := header.Header.Get("Content-Type")
	if mime == "" {
		return name, "FILE_UPLOAD_MIME_ERROR", errors.New("cant get file content-type")
	}

	// check allowed mime
	if !strings.HasPrefix(mime, Mime) {
		return name, "FILE_UPLOAD_MIME_NOT_ALLOWED", errors.New(mime + " content-type not allowed")
	}

	// copy file to desitination
	ext := getExtension(header.Filename)
	name = hash.GenerateMd5() + "." + ext
	filePath := Path + name
	out, err := os.Create(filePath)
	if err != nil {
		return name, "FILE_UPLOAD_CREATE_ERROR", err
	}

	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		return name, "FILE_UPLOAD_COPY_ERROR", err
	}

	return name, "", nil
}

func getExtension(filename string) string {
	raw := strings.Split(filename, ".")
	return raw[len(raw)-1]
}

func getFileSize(file multipart.File) (int, error) {
	type size interface {
		Size() int64
	}

	var fsize string
	var i64 int

	if sizeInterface, ok := file.(size); ok {
		sizeInterface.Size()

		fsize = fmt.Sprintf("%d", sizeInterface.Size())
	} else {
		return i64, errors.New("cant get file size")
	}

	s, err := strconv.Atoi(fsize)
	if err != nil {
		return i64, err
	}

	return int(s), nil
}

func getFileSizeLong(file multipart.File) (int64, error) {
	type size interface {
		Size() int64
	}

	var fsize string
	var i64 int64

	if sizeInterface, ok := file.(size); ok {
		sizeInterface.Size()

		fsize = fmt.Sprintf("%d", sizeInterface.Size())
	} else {
		return i64, errors.New("cant get file size")
	}

	s, err := strconv.Atoi(fsize)
	if err != nil {
		return i64, err
	}

	return int64(s), nil
}
