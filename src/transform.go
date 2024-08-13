package main

import "math"

type translation struct {
	vertical, horizontal float64
}

func (t1 *translation) equalTranslation(t2 translation) bool {
	return t1.horizontal == t2.horizontal && t1.vertical == t2.vertical
}

func findTranslation(main, secondary []Triangle) translation {
	mainTri := main[0]
	smallestDifference := math.Inf(1)
	var closestTri Triangle
	var difference float64

	for _, secondTri := range secondary {
		difference = mainTri.smallestDifference(secondTri)
		if difference < smallestDifference {
			closestTri = secondTri
			smallestDifference = difference
		}
	}

	distanceX := closestTri.c.X - mainTri.c.X
	distanceY := closestTri.c.Y - mainTri.c.Y

	return translation{distanceX, distanceY}
}
