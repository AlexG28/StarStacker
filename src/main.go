package main

import (
	"errors"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"strings"
	"sync"
)

func main() {
	argsWithProg := os.Args[1:]

	if len(argsWithProg) != 1 {
		log.Fatal("Please enter only 1 directory")
	}

	imageFiles := collectImageFiles(argsWithProg[0])
	numberOfFiles := len(imageFiles)

	triangulations := make([][]Triangle, numberOfFiles)
	images := make([]image.Image, numberOfFiles)

	var wg sync.WaitGroup
	for i, filePath := range imageFiles {
		wg.Add(1)
		go func(i int, filePath string) {
			defer wg.Done()
			triangulation, image := openFile(filePath)
			triangulations[i] = triangulation
			images[i] = image
		}(i, filePath)
	}
	wg.Wait()

	referenceTriangulation := triangulations[0]
	translations := make([]translation, numberOfFiles-1)

	for i := 1; i < numberOfFiles; i++ {
		currTranslation := findTranslation(referenceTriangulation, triangulations[i])
		translations[i-1] = currTranslation
	}

	stackedImage := images[0]

	for i := 1; i < numberOfFiles; i++ {
		stackedImage = stack(translations[i-1], stackedImage, images[i])
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

func collectImageFiles(dirPath string) []string {
	dir, err := os.Open(dirPath)

	imageList := make([]string, 0)

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
				imageList = append(imageList, fmt.Sprintf("%s/%s", dirPath, file.Name()))
			}
		}
	}

	index, err := findBaseImage(imageList)

	if err != nil {
		log.Fatalf(err.Error())
	}

	imageList[0], imageList[index] = imageList[index], imageList[0]

	return imageList
}

func findBaseImage(imageNames []string) (int, error) {
	for i, img := range imageNames {
		splitString := strings.Split(img, "/")
		lastElement := splitString[len(splitString)-1]

		if strings.HasPrefix(lastElement, "base") {
			return i, nil
		}
	}
	return -1, errors.New("no base image specified. Please prefix an image file of your choise with 'base'")
}
