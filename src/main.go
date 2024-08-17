package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"strings"
)

func main() {
	argsWithProg := os.Args[1:]

	if len(argsWithProg) != 1 {
		log.Fatal("Please enter only 1 directory")
	}

	imageFiles := process(argsWithProg[0])
	numberOfFiles := len(imageFiles)

	trianglesArray := make([][]Triangle, numberOfFiles)
	imageArray := make([]image.Image, numberOfFiles)

	for i, filePath := range imageFiles {
		triangles, image := openFile(filePath)

		trianglesArray[i] = triangles
		imageArray[i] = image
	}

	baseTriangles := trianglesArray[0]
	translations := make([]translation, numberOfFiles-1)

	for i := 1; i < numberOfFiles; i++ {
		currTranslation := findTranslation(baseTriangles, trianglesArray[i])
		translations[i-1] = currTranslation
	}

	stackedImage := imageArray[0]

	for i := 1; i < numberOfFiles; i++ {
		stackedImage = stack(translations[i-1], stackedImage, imageArray[i])
	}

	saveOutputImage(stackedImage, "output")

	fmt.Println("Done!")
}

func openFile(filePath string) ([]Triangle, image.Image) {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Printf("Unable to open file %v", filePath)
		os.Exit(1)
	}

	defer file.Close()

	img, err := png.Decode(file)

	if err != nil {
		fmt.Printf("Unable to decode file %v", filePath)
		os.Exit(1)
	}

	binaryImage := toBinary(&img, 150)
	stars := countStars(binaryImage)
	vertices := starsToVertices(*stars)
	triangles := triangulate(vertices)

	return triangles, img
}

func process(dirPath string) []string {
	dir, err := os.Open(dirPath)

	output := make([]string, 0)

	if err != nil {
		log.Fatalf("Failed to open directory: %v", err)
	}
	defer dir.Close()

	files, err := dir.ReadDir(-1)

	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	for _, file := range files {
		if !file.IsDir() {
			if strings.HasSuffix(file.Name(), ".png") {
				output = append(output, fmt.Sprintf("%s/%s", dirPath, file.Name()))
			}
		}
	}
	return output
}
