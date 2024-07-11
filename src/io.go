package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

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

func saveStarPoints(filename string, stars *[]star) error {
	filepath := fmt.Sprintf("/home/alexlinux/projects/StarCounter/testfiles/%s.txt", filename)

	file, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	for _, star := range *stars {
		// _, err = file.WriteString(star.print())
		_, err = file.WriteString(star.print())

		if err != nil {
			return err
		}
	}

	file.WriteString("\n\n\n\n")

	return nil
}
