package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func main() {
	argsWithProg := os.Args[1:]

	if len(argsWithProg) != 2 {
		fmt.Println("Please enter 2 filenames")
		os.Exit(1)
	}

	triangles1, image1 := openFile(argsWithProg[0])
	triangles2, image2 := openFile(argsWithProg[1])

	translation := findTranslation(triangles1, triangles2)

	stackedImage := stack(translation, image1, image2)

	saveOutputImage(stackedImage, "output")
}

func openFile(filename string) ([]Triangle, image.Image) {
	filepath := fmt.Sprintf("/home/alexlinux/projects/StarCounter/testfiles/%s.png", filename)

	file, err := os.Open(filepath)

	if err != nil {
		fmt.Printf("Unable to open file %v", filename)
		os.Exit(1)
	}

	defer file.Close()

	img, err := png.Decode(file)

	if err != nil {
		fmt.Printf("Unable to decode file %v", filename)
		os.Exit(1)
	}

	binaryImage := toBinary(&img, 150)
	stars := countStars(binaryImage)
	vertices := starsToVertices(*stars)
	triangles := triangulate(vertices)

	return triangles, img
}
