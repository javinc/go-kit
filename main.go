package main

import (
	"image/jpeg"
	"log"
	"os"

	"github.com/javinc/go-kit/file/image"
	"github.com/nfnt/resize"
)

func main() {
	// baseTest()
	newPath, err := image.Thumbnail("test.jpg", 600, 300)
	log.Println(newPath)
	log.Println(err)
}

func baseTest() {
	// open "test.jpg"
	file, err := os.Open("test.jpg")
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// log.Println(img)

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	// m := resize.Resize(900, 0, img, resize.Lanczos3)
	m := resize.Thumbnail(600, 600, img, resize.Lanczos3)
	out, err := os.Create("test_resized2.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
}
