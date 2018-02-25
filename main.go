package main

import "github.com/javinc/go-kit/file/image"

func main() {
	image.Thumbnail("test.jpg", 600, 400)
}
