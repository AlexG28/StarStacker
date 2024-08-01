package main

import (
	"math"
	"testing"
)

func TestTriangulate(t *testing.T) {
	stars := make([]star, 3)

	stars[0] = star{
		id: 1,
		points: [][2]int{
			{10, 10},
		},
		location: [2]int{10, 10},
	}
	stars[1] = star{
		id: 2,
		points: [][2]int{
			{20, 20},
		},
		location: [2]int{20, 20},
	}
	stars[2] = star{
		id: 3,
		points: [][2]int{
			{10, 20},
		},
		location: [2]int{10, 20},
	}

	expected := Triangle{
		Vertex{15, 1020},
		Vertex{-990, -990},
		Vertex{1020, -990},
	}

	result := superTriangle(stars)

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
