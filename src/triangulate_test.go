package main

import (
	"math"
	"testing"
)

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

func TestUniqueEdge(t *testing.T) {
	badTriangles := make([]Triangle, 3)

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

	res := uniqueEdges(badTriangles)

	expectedUniques := 7

	if len(res) != expectedUniques {
		t.Errorf("Expected %v uniques but got %v", expectedUniques, len(res))
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

	triangles[2] = Triangle{
		Vertex{2, 1},
		Vertex{4, 4},
		Vertex{3, 8},
	}

	st := Triangle{
		Vertex{3, 8},
		Vertex{-10, -10},
		Vertex{10, -10},
	}

	res := removeSuperTriangle(triangles, st)
	expected := 3

	if len(res) != expected {
		t.Errorf("Got a length of %v but expecting %v", len(res), expected)
	}
}

// func TestTriangulation(t *testing.T) {
// 	stars := make([]Vertex, 5)

// 	stars[0] = Vertex{0, 0}
// 	stars[1] = Vertex{1, 0}
// 	stars[2] = Vertex{0, 1}
// 	stars[3] = Vertex{1, 1}
// 	stars[4] = Vertex{2, 2}

// 	result := triangulate(stars)

// 	if len(result) == 0 {
// 		t.Errorf("It didn't fucking work :(")
// 	}
// }
