package main

import (
	"image"
	"image/color"
)

func stack(trans translation, baseImg, sideImg image.Image) image.Image {
	baseImageSize := baseImg.Bounds()
	sideImgSize := sideImg.Bounds()

	sideImgSize.Min.X = sideImgSize.Min.X + int(trans.horizontal)
	sideImgSize.Max.X = sideImgSize.Max.X + int(trans.horizontal)

	sideImgSize.Min.Y = sideImgSize.Min.Y + int(trans.vertical)
	sideImgSize.Max.Y = sideImgSize.Max.Y + int(trans.vertical)

	resultBounds := baseImageSize.Intersect(sideImgSize)
	output := image.NewRGBA(resultBounds)

	for y := resultBounds.Min.Y; y < resultBounds.Max.Y; y++ {
		for x := resultBounds.Min.X; x < resultBounds.Max.X; x++ {
			c1 := baseImg.At(x-int(trans.horizontal), y-int(trans.vertical))
			c2 := sideImg.At(x, y)

			r1, g1, b1, a1 := c1.RGBA()
			r2, g2, b2, a2 := c2.RGBA()

			r := uint8((r1/257 + r2/257) / 2)
			g := uint8((g1/257 + g2/257) / 2)
			b := uint8((b1/257 + b2/257) / 2)
			a := uint8((a1/257 + a2/257) / 2)

			output.Set(x, y, color.RGBA{r, g, b, a})

		}
	}

	return output
}
