package main

import (
	"image"
	"image/color"
)

func toBinary(img *image.Image, threshold uint8) *image.Gray {
	bounds := (*img).Bounds()
	binaryImage := image.NewGray(bounds)

	white := color.Gray{Y: 255}
	black := color.Gray{Y: 0}

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := (*img).At(x, y)
			grayPixel := color.GrayModel.Convert(pixel).(color.Gray)

			if grayPixel.Y >= threshold {
				binaryImage.SetGray(x, y, white)
			} else {
				binaryImage.SetGray(x, y, black)
			}
		}
	}

	return binaryImage
}
