package main

import (
	"math"
	"math/rand/v2"
)

type translation struct {
	vertical, horizontal float64
}

func (t1 *translation) equalTranslation(t2 translation) bool {
	return t1.horizontal == t2.horizontal && t1.vertical == t2.vertical
}

func findTranslation(main, secondary []Triangle) translation {
	i1, i2, i3 := generate3RandomUniqueTriangleIndices(len(main))

	t1, t2, t3 := main[i1], main[i2], main[i3]

	dist1, tri1 := findSmallestDifferenceBetweenTriangleAndTriangulation(t1, secondary)
	dist2, tri2 := findSmallestDifferenceBetweenTriangleAndTriangulation(t2, secondary)
	dist3, tri3 := findSmallestDifferenceBetweenTriangleAndTriangulation(t3, secondary)

	var translation translation

	if dist1 <= dist2 && dist1 <= dist3 {
		translation = translationBetweenTwoTriangles(tri1, t1)
	} else if dist2 <= dist1 && dist2 <= dist3 {
		translation = translationBetweenTwoTriangles(tri2, t2)
	} else {
		translation = translationBetweenTwoTriangles(tri3, t3)
	}

	return translation
}

func findSmallestDifferenceBetweenTriangleAndTriangulation(base Triangle, secondary []Triangle) (float64, Triangle) {
	smallestDifference := math.Inf(1)
	var closestTri Triangle
	var difference float64

	for _, secondTri := range secondary {
		difference = base.smallestDifference(secondTri)
		if difference < smallestDifference {
			closestTri = secondTri
			smallestDifference = difference
		}

		if difference == 0 {
			break
		}
	}

	return smallestDifference, closestTri
}

func translationBetweenTwoTriangles(otherTri, baseTri Triangle) translation {
	distanceX := otherTri.c.X - baseTri.c.X
	distanceY := otherTri.c.Y - baseTri.c.Y
	return translation{math.Round(distanceY), math.Round(distanceX)}
}

func generate3RandomUniqueTriangleIndices(mainLength int) (int, int, int) {
	if mainLength <= 3 {
		if mainLength == 3 {
			return 0, 1, 2
		} else if mainLength == 2 {
			return 0, 1, 0
		} else {
			return 0, 0, 0
		}
	}

	uniqueNumbers := make(map[int32]bool)
	var nums []int

	for len(uniqueNumbers) < 3 {
		num := rand.Int32N(int32(mainLength))
		uniqueNumbers[num] = true
	}

	for num := range uniqueNumbers {
		nums = append(nums, int(num))
	}

	return nums[0], nums[1], nums[2]
}
