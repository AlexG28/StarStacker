package main

import (
	"testing"
)

func TestTriangularize(t *testing.T) {
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
