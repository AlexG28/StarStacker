package main

import (
	"fmt"
	"image/jpeg"
	"os"
)

func main() {
	argsWithProg := os.Args[1:]
	fmt.Printf("the filenames are: %v\n", argsWithProg)

	filename := argsWithProg[0]
	filepath := fmt.Sprintf("/home/alexlinux/projects/StarCounter/testfiles/%s.jpg", filename)

	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println("unable to open jpeg")
		os.Exit(1)
	}

	defer file.Close()

	img, err := jpeg.Decode(file)

	if err != nil {
		fmt.Println("Unable to decode jpeg")
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
