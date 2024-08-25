package main

import (
	"image"
	"image/color"
)

func stack(trans translation, baseImg, sideImg image.Image) image.Image {
	baseImageSize := baseImg.Bounds()
	sideImgSize := sideImg.Bounds()

	resultBounds := baseImageSize.Intersect(sideImgSize)
	output := image.NewRGBA(resultBounds)

	for y := resultBounds.Min.Y; y < resultBounds.Max.Y; y++ {
		for x := resultBounds.Min.X; x < resultBounds.Max.X; x++ {
			sideImgX, sideImgY := x+int(trans.horizontal), y+int(trans.vertical)

			c1 := baseImg.At(x, y)
			c2 := sideImg.At(sideImgX, sideImgY)

			r1, g1, b1, _ := c1.RGBA()
			r2, g2, b2, _ := c2.RGBA()

			r := uint8((r1/257 + r2/257) / 2)
			g := uint8((g1/257 + g2/257) / 2)
			b := uint8((b1/257 + b2/257) / 2)
			a := uint8(255)

			output.Set(x, y, color.RGBA{r, g, b, a})
		}
	}
	return output
}
