package main

import (
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

func TestBasicGetAdjacentPoints(t *testing.T) {
	visited := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	x, y := 1, 1

	result := getAdjacentPoints(x, y, &visited)

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
	visited := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	x, y := 1, 0

	result := getAdjacentPoints(x, y, &visited)

	/*
		NOTE: the x position (first elem) in the pars is the vertical selector (outer) while the y value is the inner (horizontal)
	*/

	expected := [][2]int{
		{0, 0},
		{0, 1},
		{1, 1},
		{2, 0},
		{2, 1},
	}

	if !sameAdjacentPoints(expected, result) {
		t.Errorf("resulted: %v    expected: %v", result, expected)
	}
}

func TestBottomSideGetAdjacentPoints(t *testing.T) {
	visited := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	x, y := 2, 1

	result := getAdjacentPoints(x, y, &visited)

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
	visited := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	x, y := 2, 2

	result := getAdjacentPoints(x, y, &visited)

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
	visited := [][]int{
		{0, 0, 1},
		{0, 0, 2},
		{0, 0, 3},
	}

	x, y := 1, 1

	result := getAdjacentPoints(x, y, &visited)

	expected := [][2]int{
		{0, 0},
		{0, 1},
		{1, 0},
		{2, 0},
		{2, 1},
	}

	if !sameAdjacentPoints(expected, result) {
		t.Errorf("resulted: %v    expected: %v", result, expected)
	}
}

func TestGetAdjacentPointsWithVisitedInCorner(t *testing.T) {
	visited := [][]int{
		{0, 0, 0},
		{0, 14, 0},
		{0, 0, 0},
	}

	x, y := 2, 2

	result := getAdjacentPoints(x, y, &visited)

	expected := [][2]int{
		{2, 1},
		{1, 2},
	}

	if !sameAdjacentPoints(expected, result) {
		t.Errorf("resulted: %v    expected: %v", result, expected)
	}
}

func TestGetAdjacentPointsWithNone(t *testing.T) {
	visited := [][]int{
		{0, 0, 0},
		{0, 14, 2},
		{0, 11, 0},
	}

	x, y := 2, 2

	result := getAdjacentPoints(x, y, &visited)

	expected := [][2]int{}

	if !sameAdjacentPoints(expected, result) {
		t.Errorf("resulted: %v    expected: %v", result, expected)
	}
}
