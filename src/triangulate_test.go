package main

import (
	"math"
	"testing"
)

func triangleToMap(t []Triangle) map[string]struct{} {
	hashes := make(map[string]struct{})
	for _, tri := range t {
		hash := tri.Hash()
		hashes[hash] = struct{}{}
	}
	return hashes
}

func sameArrayOFTriangles(t1, t2 []Triangle) bool {
	len1 := len(t1)
	len2 := len(t2)

	if len1 != len2 {
		return false
	}

	t1dict := triangleToMap(t1)
	t2dict := triangleToMap(t2)

	for hash, _ := range t1dict {
		if _, ok := t2dict[hash]; !ok {
			return false
		}
	}

	for hash, _ := range t2dict {
		if _, ok := t1dict[hash]; !ok {
			return false
		}
	}

	return true
}

func TestTriangulate(t *testing.T) {
	expected := Triangle{
		Vertex{15, 1020},
		Vertex{-990, -990},
		Vertex{1020, -990},
	}

	stars := [3]Vertex{
		{10, 10},
		{20, 20},
		{10, 20},
	}
	starSlice := stars[:]

	result := superTriangle(starSlice)

	if expected != result {
		t.Errorf("result: %v      expected: %v", result, expected)
	}

}

func TestCircumCircle(t *testing.T) {
	epsilon := 1e-6
	tri := Triangle{
		Vertex{3, 4},
		Vertex{5, 9},
		Vertex{4, 8},
	}

	center, radius := circumCenter(tri)

	expectedCenter := Vertex{8.166667, 4.833333}
	expectedRadius := 5.233439

	if math.Abs(center.X-expectedCenter.X) > epsilon || math.Abs(center.Y-expectedCenter.Y) > epsilon {
		t.Errorf("circumcenters not matched: %v      expected: %v", center, expectedCenter)
		t.Errorf("The difference is   X: %v   Y: %v", math.Abs(center.X-expectedCenter.X), math.Abs(center.Y-expectedCenter.Y))
	}

	if math.Abs(radius-expectedRadius) > epsilon {
		t.Errorf("radius not matched: %v      expected: %v", radius, expectedRadius)
		t.Errorf("The difference is:   %v", math.Abs(radius-expectedRadius))
	}
}

func TestCircumCircleSmall(t *testing.T) {
	epsilon := 1e-6
	degenerateTri := Triangle{
		Vertex{5.567926, 5.567156},
		Vertex{5.567122, 5.567826},
		Vertex{5.567555, 5.567000},
	}

	center, radius := circumCenter(degenerateTri)

	expectedCenter := Vertex{5.56755247276, 5.56752516732}
	expectedRadius := 0.000525173404322

	if math.Abs(center.X-expectedCenter.X) > epsilon || math.Abs(center.Y-expectedCenter.Y) > epsilon {
		t.Errorf("circumcenters not matched: %v      expected: %v", center, expectedCenter)
		t.Errorf("The difference is   X: %v   Y: %v", math.Abs(center.X-expectedCenter.X), math.Abs(center.Y-expectedCenter.Y))
	}

	if math.Abs(radius-expectedRadius) > epsilon {
		t.Errorf("radius not matched: %v      expected: %v", radius, expectedRadius)
		t.Errorf("The difference is:   %v", math.Abs(radius-expectedRadius))
	}
}

func TestCircumCircleLarge(t *testing.T) {
	epsilon := 1e-6
	degenerateTri := Triangle{
		Vertex{187, 41},
		Vertex{557, 666},
		Vertex{972, 420},
	}

	center, radius := circumCenter(degenerateTri)

	expectedCenter := Vertex{579.391836356, 230.724032877}
	expectedRadius := 435.851536524

	if math.Abs(center.X-expectedCenter.X) > epsilon || math.Abs(center.Y-expectedCenter.Y) > epsilon {
		t.Errorf("circumcenters not matched: %v      expected: %v", center, expectedCenter)
		t.Errorf("The difference is   X: %v   Y: %v", math.Abs(center.X-expectedCenter.X), math.Abs(center.Y-expectedCenter.Y))
	}

	if math.Abs(radius-expectedRadius) > epsilon {
		t.Errorf("radius not matched: %v      expected: %v", radius, expectedRadius)
		t.Errorf("The difference is:   %v", math.Abs(radius-expectedRadius))
	}
}

func TestInCircumcircle(t *testing.T) {
	tri := Triangle{
		Vertex{3, 4},
		Vertex{5, 9},
		Vertex{4, 8},
	}
	point := Vertex{10, 20}

	res := inCircumcircle(tri, point)
	expected := false

	if res != expected {
		t.Errorf("Incorrect inCircumcircle: %v    expected: %v", res, expected)
	}
}

func TestBoundaryOfPolygonalHole(t *testing.T) {
	badTriangles := make([]Triangle, 4)

	badTriangles[0] = Triangle{
		Vertex{0, 0},
		Vertex{1, 1},
		Vertex{1, 2},
	}
	badTriangles[1] = Triangle{
		Vertex{1, 1},
		Vertex{1, 2},
		Vertex{2, 1},
	}
	badTriangles[2] = Triangle{
		Vertex{1, 2},
		Vertex{2, 1},
		Vertex{4, 4},
	}
	badTriangles[3] = Triangle{
		Vertex{4, 4},
		Vertex{2, 1},
		Vertex{7, 1},
	}

	res := boundaryOfPolygonalHole(badTriangles)

	expectedUniques := 6

	if len(res) != expectedUniques {
		t.Errorf("Expected %v uniques but got %v", expectedUniques, len(res))
	}
}

func TestRemoveBadTriangles(t *testing.T) {
	triangulation := make([]Triangle, 4)

	triangulation[0] = Triangle{
		Vertex{0, 0},
		Vertex{1, 1},
		Vertex{1, 2},
	}
	triangulation[1] = Triangle{
		Vertex{1, 1},
		Vertex{1, 2},
		Vertex{2, 1},
	}
	triangulation[2] = Triangle{
		Vertex{1, 2},
		Vertex{2, 1},
		Vertex{4, 4},
	}
	triangulation[3] = Triangle{
		Vertex{3, 0},
		Vertex{0, 0},
		Vertex{1, 1},
	}

	badTriangles := make([]Triangle, 2)

	badTriangles[0] = Triangle{
		Vertex{0, 0},
		Vertex{1, 1},
		Vertex{1, 2},
	}
	badTriangles[1] = Triangle{
		Vertex{1, 1},
		Vertex{1, 2},
		Vertex{2, 1},
	}

	result := removeBadTrianglesFromTriangulation(triangulation, badTriangles)

	expected := []Triangle{
		{Vertex{1, 2}, Vertex{2, 1}, Vertex{4, 4}},
		{Vertex{3, 0}, Vertex{0, 0}, Vertex{1, 1}},
	}

	if !sameArrayOFTriangles(result, expected) {
		t.Errorf("Expected %v uniques but got %v", expected, result)
	}
}

func TestRemoveSuperTriangle(t *testing.T) {
	triangles := make([]Triangle, 4)

	triangles[0] = Triangle{
		Vertex{0, 0},
		Vertex{1, 1},
		Vertex{1, 2},
	}
	triangles[1] = Triangle{
		Vertex{1, 1},
		Vertex{1, 2},
		Vertex{2, 1},
	}
	triangles[2] = Triangle{
		Vertex{1, 2},
		Vertex{2, 1},
		Vertex{4, 4},
	}

	triangles[3] = Triangle{
		Vertex{2, 1},
		Vertex{4, 4},
		Vertex{3, 8},
	}

	st := Triangle{
		Vertex{3, 8},
		Vertex{-10, -10},
		Vertex{10, -10},
	}

	result := removeSuperTriangle(triangles, st)
	expected := []Triangle{
		{Vertex{0, 0}, Vertex{1, 1}, Vertex{1, 2}},
		{Vertex{1, 1}, Vertex{1, 2}, Vertex{2, 1}},
		{Vertex{1, 2}, Vertex{2, 1}, Vertex{4, 4}},
	}

	if !sameArrayOFTriangles(result, expected) {
		t.Errorf("Got a length of %v but expecting %v", result, expected)
	}
}

func TestTriangulationBasic(t *testing.T) {
	stars := make([]Vertex, 3)

	stars[0] = Vertex{0, 0}
	stars[1] = Vertex{5, 2}
	stars[2] = Vertex{2, 5}

	result := triangulate(stars)
	expected := []Triangle{
		{Vertex{0, 0}, Vertex{5, 2}, Vertex{2, 5}},
	}

	if !sameArrayOFTriangles(result, expected) {
		t.Errorf("Result: %v   expected: %v", result, expected)
	}
}

func TestTriangulationMedium(t *testing.T) {
	stars := make([]Vertex, 4)

	stars[0] = Vertex{0, 0}
	stars[1] = Vertex{-10, 8}
	stars[2] = Vertex{4, 1}
	stars[3] = Vertex{5, 17}

	result := triangulate(stars)
	expected := []Triangle{
		{Vertex{-10, 8}, Vertex{0, 0}, Vertex{5, 17}},
		{Vertex{0, 0}, Vertex{4, 1}, Vertex{5, 17}},
	}

	if !sameArrayOFTriangles(expected, result) {
		t.Errorf("Result: %v   expected: %v", result, expected)
	}
}

func TestTriangulationAdvanced(t *testing.T) {
	stars := make([]Vertex, 6)

	stars[0] = Vertex{0, 0}
	stars[1] = Vertex{0, 8}
	stars[2] = Vertex{4, 1}
	stars[3] = Vertex{5, 7}
	stars[4] = Vertex{8, 2}
	stars[5] = Vertex{9, 14}

	expected := []Triangle{
		{Vertex{0, 0}, Vertex{0, 8}, Vertex{4, 1}},
		{Vertex{0, 8}, Vertex{4, 1}, Vertex{5, 7}},
		{Vertex{4, 1}, Vertex{8, 2}, Vertex{5, 7}},
		{Vertex{0, 8}, Vertex{5, 7}, Vertex{9, 14}},
		{Vertex{5, 7}, Vertex{8, 2}, Vertex{9, 14}},
	}

	result := triangulate(stars)

	if !sameArrayOFTriangles(result, expected) {
		t.Errorf("Result: %v   expected: %v", result, expected)
	}
}
