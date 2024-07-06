package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func main() {
	argsWithProg := os.Args[1:]
	fmt.Printf("the filenames are: %v\n", argsWithProg)

	filename := argsWithProg[0]
	filepath := fmt.Sprintf("/home/alexlinux/projects/StarCounter/testfiles/%s.png", filename)

	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println("unable to open png")
		os.Exit(1)
	}

	defer file.Close()

	img, err := png.Decode(file)

	if err != nil {
		fmt.Println("Unable to decode png")
		os.Exit(1)
	}

	bounds := img.Bounds()

	fmt.Printf("bounds: %v\n", bounds)

	toWrite := fmt.Sprintf("The image stretches from (%d, %d) to (%d, %d)", bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y)
	err = writeOutput(toWrite, "out")

	if err != nil {
		fmt.Println("Unable to write result to file")
		os.Exit(1)
	}

	newGrayScaleImg := toGrayScale(&img)
	saveOutputImage(newGrayScaleImg, "grayscale")

	newTest := toBinary(newGrayScaleImg, 30)
	saveOutputImage(newTest, "binary")
}

func writeOutput(text string, filename string) error {
	filepath := fmt.Sprintf("/home/alexlinux/projects/StarCounter/testfiles/%s.txt", filename)
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(text)
	if err != nil {
		return err
	}

	return nil
}

func saveOutputImage(img *image.Gray, filename string) error {
	filepath := fmt.Sprintf("/home/alexlinux/projects/StarCounter/testfiles/%s.png", filename)

	file, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	err = png.Encode(file, img)

	if err != nil {
		panic(err)
	}

	return nil
}
