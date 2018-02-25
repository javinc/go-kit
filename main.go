package main

import (
	"log"

	"github.com/javinc/go-kit/file/image"
)

func main() {
	// baseTest()
	newPath, err := image.Thumbnail("test.png", 600, 300)
	log.Println(newPath)
	log.Println(err)

	image.Thumbnail("test.jpg", 600, 300)
}
