package main

import (
	"image"
	"image/color"
)

func toGrayScale(img *image.Image) image.Image {
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
