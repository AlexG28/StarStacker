package main

import (
	"image"
	"image/color"
)

func toGrayScale(img *image.Image) *image.Gray {
	bounds := (*img).Bounds()
	grayImg := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			old := (*img).At(x, y)
			newGray := color.GrayModel.Convert(old)
			grayImg.Set(x, y, newGray)
		}
	}

	return grayImg
}

func toBinary(img *image.Gray, threshold uint8) *image.Gray {
	bounds := (*img).Bounds()
	binaryImage := image.NewGray(bounds)

	white := color.Gray{Y: 255}
	black := color.Gray{Y: 0}

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			old := (*img).GrayAt(x, y)

			if old.Y >= threshold {
				binaryImage.SetGray(x, y, white)
			} else {
				binaryImage.SetGray(x, y, black)
			}
		}
	}

	return binaryImage
}
