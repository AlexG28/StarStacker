package main

import (
	"image"
	"image/color"
	"testing"
)

func sliceToMap(points [][2]int) map[[2]int]struct{} {
	pointMap := make(map[[2]int]struct{})
	for _, pair := range points {
		pointMap[pair] = struct{}{}
	}
	return pointMap
}

func sameAdjacentPoints(left, right [][2]int) bool {
	if len(left) != len(right) {
		return false
	}

	leftMap := sliceToMap(left)
	rightMap := sliceToMap(right)

	for pair := range leftMap {
		_, exists := rightMap[pair]
		if !exists {
			return false
		}
	}

	return true
}

func makeBasicGrayImage(ima *[][]int) *image.Gray {
	height := len(*ima)
	width := len((*ima)[0])

	max := image.Point{height, width}
	min := image.Point{0, 0}

	image := image.NewGray(image.Rectangle{min, max})

	for x := range height {
		for y := range width {
			brightness := int8((*ima)[x][y] * 255)
			image.SetGray(x, y, color.Gray{uint8(brightness)})
		}
	}

	return image
}

func TestBasicGetAdjacentPoints(t *testing.T) {
	image := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}

	visited := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	x, y := 1, 1

	img := makeBasicGrayImage(&image)

	result := getAdjacentPoints(x, y, &visited, img)

	expected := [][2]int{
		{0, 0},
		{0, 1},
		{0, 2},
		{1, 0},
		{1, 2},
		{2, 0},
		{2, 1},
		{2, 2},
	}

	if !sameAdjacentPoints(expected, result) {
		t.Errorf("expected: %v    resulted: %v", expected, result)
	}
}

func TestLeftSideGetAdjacentPoints(t *testing.T) {
	image := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}

	visited := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	x, y := 1, 0
	img := makeBasicGrayImage(&image)

	result := getAdjacentPoints(x, y, &visited, img)

	/*
		NOTE: the x position (first elem) in the pars is the vertical selector (outer) while the y value is the inner (horizontal)
	*/

	expected := [][2]int{
		{0, 0},
		{0, 1},
		{2, 0},
		{2, 1},
	}

	if !sameAdjacentPoints(expected, result) {
		t.Errorf("resulted: %v    expected: %v", result, expected)
	}
}

func TestBottomSideGetAdjacentPoints(t *testing.T) {
	image := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}

	visited := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	x, y := 2, 1
	img := makeBasicGrayImage(&image)
	result := getAdjacentPoints(x, y, &visited, img)

	expected := [][2]int{
		{2, 0},
		{1, 0},
		{1, 1},
		{1, 2},
		{2, 2},
	}

	if !sameAdjacentPoints(expected, result) {
		t.Errorf("resulted: %v    expected: %v", result, expected)
	}
}

func TestCornerSideGetAdjacentPoints(t *testing.T) {
	image := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}

	visited := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	x, y := 2, 2
	img := makeBasicGrayImage(&image)
	result := getAdjacentPoints(x, y, &visited, img)

	expected := [][2]int{
		{2, 1},
		{1, 1},
		{1, 2},
	}

	if !sameAdjacentPoints(expected, result) {
		t.Errorf("resulted: %v    expected: %v", result, expected)
	}
}

func TestGetAdjacentPointsWithVisited(t *testing.T) {
	image := [][]int{
		{0, 0, 1},
		{1, 1, 1},
		{1, 1, 1},
	}

	visited := [][]int{
		{0, 0, 1},
		{0, 0, 2},
		{0, 0, 3},
	}

	x, y := 1, 1
	img := makeBasicGrayImage(&image)
	result := getAdjacentPoints(x, y, &visited, img)

	expected := [][2]int{
		{1, 0},
		{2, 0},
		{2, 1},
	}

	if !sameAdjacentPoints(expected, result) {
		t.Errorf("resulted: %v    expected: %v", result, expected)
	}
}

func TestGetAdjacentPointsWithVisitedInCorner(t *testing.T) {
	image := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}

	visited := [][]int{
		{0, 0, 0},
		{0, 14, 0},
		{0, 0, 0},
	}

	x, y := 2, 2
	img := makeBasicGrayImage(&image)
	result := getAdjacentPoints(x, y, &visited, img)

	expected := [][2]int{
		{2, 1},
		{1, 2},
	}

	if !sameAdjacentPoints(expected, result) {
		t.Errorf("resulted: %v    expected: %v", result, expected)
	}
}

func TestGetAdjacentPointsWithNone(t *testing.T) {
	image := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}

	visited := [][]int{
		{0, 0, 0},
		{0, 14, 2},
		{0, 11, 0},
	}

	x, y := 2, 2
	img := makeBasicGrayImage(&image)
	result := getAdjacentPoints(x, y, &visited, img)

	expected := [][2]int{}

	if !sameAdjacentPoints(expected, result) {
		t.Errorf("resulted: %v    expected: %v", result, expected)
	}
}

// func TestBFSBasic(t *testing.T) {
// 	visited := [][]int{
// 		{1, 0, 0, 0, 0},
// 		{1, 1, 0, 0, 0},
// 		{0, 1, 1, 0, 0},
// 		{0, 0, 0, 0, 0},
// 		{0, 0, 0, 0, 0},
// 	}

// 	x, y := 0, 0

// 	/*
// 		getAdjacentPoints needs to also look at the GrayMap values
// 		so does BFS
// 		this is clearly not ready yet

// 	*/

// 	result := bfs(x, y, &visited)

// 	expected := [][2]int{
// 		{0, 0},
// 		{1, 0},
// 		{1, 1},
// 		{2, 1},
// 		{2, 2},
// 	}

// 	if !sameAdjacentPoints(expected, result) {
// 		t.Errorf("resulted: %v    expected: %v", result, expected)
// 	}
// }
