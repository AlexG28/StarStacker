package main

import (
	"fmt"
	"image/png"
	"os"
)

func main() {
	argsWithProg := os.Args[1:]
	filename := argsWithProg[0]
	filepath := fmt.Sprintf("/home/alexlinux/projects/StarCounter/testfiles/%s.png", filename)

	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println("Unable to open file")
		os.Exit(1)
	}

	defer file.Close()

	img, err := png.Decode(file)

	if err != nil {
		fmt.Println("Unable to decode file")
		os.Exit(1)
	}

	binaryImage := toBinary(&img, 30)
	saveOutputImage(binaryImage, "binary")

	stars := countStars(binaryImage)
	count := len(*stars)
	fmt.Printf("count: %v\n", count)

	// triangles := triangularize(*stars)

	saveStarPoints("stars", stars)
}
