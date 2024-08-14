package main

import (
	"image"
	"image/color"
	"testing"
)

func TestBasicMerge(t *testing.T) {

	img1 := image.NewRGBA(image.Rect(0, 0, 5, 5))
	img2 := image.NewRGBA(image.Rect(0, 0, 5, 5))

	img1.Set(0, 0, color.RGBA{255, 0, 0, 255})
	img1.Set(0, 1, color.RGBA{255, 0, 0, 255})
	img1.Set(1, 0, color.RGBA{255, 0, 0, 255})
	img1.Set(1, 1, color.RGBA{255, 0, 0, 255})
	img1.Set(1, 2, color.RGBA{255, 0, 0, 255})
	img1.Set(2, 0, color.RGBA{255, 0, 0, 255})
	img1.Set(2, 1, color.RGBA{255, 0, 0, 255})

	img2.Set(2, 1, color.RGBA{0, 255, 255, 255})
	img2.Set(2, 2, color.RGBA{0, 255, 255, 255})
	img2.Set(3, 1, color.RGBA{0, 255, 255, 255})
	img2.Set(3, 2, color.RGBA{0, 255, 255, 255})
	img2.Set(3, 3, color.RGBA{0, 255, 255, 255})
	img2.Set(4, 1, color.RGBA{0, 255, 255, 255})
	img2.Set(4, 2, color.RGBA{0, 255, 255, 255})

	transformation1 := translation{1, 2}

	res := stack(transformation1, img1, img2)

	sizes := res.Bounds()

	whiteColours := 0
	for y := sizes.Min.Y; y < sizes.Max.Y; y++ {
		for x := sizes.Min.X; x < sizes.Max.X; x++ {
			colour := res.At(x, y)
			r, g, b, _ := colour.RGBA()
			r = r / 257
			g = g / 257
			b = b / 257

			if r == 127 && g == 127 && b == 127 {
				whiteColours += 1
			}
		}
	}

	expected := 7

	if whiteColours != expected {
		t.Errorf("Expected: %v    Got: %v", expected, whiteColours)
	}

}
