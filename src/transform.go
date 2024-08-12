package main

import "math"

type translation struct {
	vertical, horizontal float32
}

func findTranslation(main, secondary []Triangle) translation {
	/*
		take first, middle, last triangles from secondary, and find their closest triangles in main
	*/

	tri := main[0]
	currentLowest := math.MaxFloat32

	for _, secondTri := range secondary {
		new := tri.smallestDifference(secondTri)
		currentLowest = min(new, currentLowest)
	}

	translate := translation{0.0, 0.0}
	return translate

}
