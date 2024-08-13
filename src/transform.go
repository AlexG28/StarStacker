package main

import "math"

type translation struct {
	vertical, horizontal float64
}

func (t1 *translation) equalTranslation(t2 translation) bool {
	return t1.horizontal == t2.horizontal && t1.vertical == t2.vertical
}

func findTranslation(main, secondary []Triangle) translation {
	tri := main[0]
	currentLowest := math.MaxFloat32
	var currClosestTri Triangle
	var new float64

	for _, secondTri := range secondary {
		new = tri.smallestDifference(secondTri)
		if new < currentLowest {
			currClosestTri = secondTri
			currentLowest = new
		}
	}

	distanceX := currClosestTri.c.X - tri.c.X
	distanceY := currClosestTri.c.Y - tri.c.Y

	translate := translation{distanceX, distanceY}
	return translate

}
