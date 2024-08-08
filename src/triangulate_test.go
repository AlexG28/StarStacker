package main

import (
	"math"
	"math/rand/v2"
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

	for hash := range t1dict {
		if _, ok := t2dict[hash]; !ok {
			return false
		}
	}

	for hash := range t2dict {
		if _, ok := t1dict[hash]; !ok {
			return false
		}
	}

	return true
}

func generateRandomPoints(n int) []Vertex {
	points := make([]Vertex, n)

	for i := 0; i < n; i++ {
		points[i] = Vertex{
			X: rand.Float64() * 10000,
			Y: rand.Float64() * 10000,
		}
	}

	return points
}

func arrayOfTrianglesToMap(triangles []Triangle) map[string]Triangle {
	out := make(map[string]Triangle)
	for _, tri := range triangles {
		out[tri.Hash()] = tri
	}
	return out
}

func mapOfTrianglesToArray(trimap map[string]Triangle) []Triangle {
	out := make([]Triangle, 0)
	for _, tri := range trimap {
		out = append(out, tri)
	}
	return out
}

func TestTriangulate(t *testing.T) {
	expected := NewTriangle(
		Vertex{15, 1020},
		Vertex{-990, -990},
		Vertex{1020, -990},
	)

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
	center, radius := circumCenter(Vertex{3, 4}, Vertex{5, 9}, Vertex{4, 8})

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
	center, radius := circumCenter(Vertex{5.567926, 5.567156},
		Vertex{5.567122, 5.567826},
		Vertex{5.567555, 5.567000})

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
	center, radius := circumCenter(Vertex{187, 41},
		Vertex{557, 666},
		Vertex{972, 420})

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
	tri := NewTriangle(
		Vertex{3, 4},
		Vertex{5, 9},
		Vertex{4, 8},
	)
	point := Vertex{10, 20}

	res := inCircumcircle(tri, point)
	expected := false

	if res != expected {
		t.Errorf("Incorrect inCircumcircle: %v    expected: %v", res, expected)
	}
}

func TestBoundaryOfPolygonalHole(t *testing.T) {
	badTriangles := []Triangle{
		NewTriangle(Vertex{0, 0}, Vertex{1, 1}, Vertex{1, 2}),
		NewTriangle(Vertex{1, 1}, Vertex{1, 2}, Vertex{2, 1}),
		NewTriangle(Vertex{1, 2}, Vertex{2, 1}, Vertex{4, 4}),
		NewTriangle(Vertex{4, 4}, Vertex{2, 1}, Vertex{7, 1}),
	}

	res := boundaryOfPolygonalHole(arrayOfTrianglesToMap(badTriangles))

	expectedUniques := 6

	if len(res) != expectedUniques {
		t.Errorf("Expected %v uniques but got %v", expectedUniques, len(res))
	}
}

func TestRemoveBadTriangles(t *testing.T) {
	triangulation := []Triangle{
		NewTriangle(Vertex{0, 0}, Vertex{1, 1}, Vertex{1, 2}),
		NewTriangle(Vertex{1, 1}, Vertex{1, 2}, Vertex{2, 1}),
		NewTriangle(Vertex{1, 2}, Vertex{2, 1}, Vertex{4, 4}),
		NewTriangle(Vertex{3, 0}, Vertex{0, 0}, Vertex{1, 1}),
	}

	badTriangles := []Triangle{
		NewTriangle(Vertex{0, 0}, Vertex{1, 1}, Vertex{1, 2}),
		NewTriangle(Vertex{1, 1}, Vertex{1, 2}, Vertex{2, 1}),
	}

	result := removeBadTrianglesFromTriangulation(
		arrayOfTrianglesToMap(triangulation), arrayOfTrianglesToMap(badTriangles))

	expected := []Triangle{
		NewTriangle(Vertex{1, 2}, Vertex{2, 1}, Vertex{4, 4}),
		NewTriangle(Vertex{3, 0}, Vertex{0, 0}, Vertex{1, 1}),
	}

	if !sameArrayOFTriangles(mapOfTrianglesToArray(result), expected) {
		t.Errorf("Expected %v uniques but got %v", expected, result)
	}
}

func TestRemoveSuperTriangle(t *testing.T) {
	triangles := []Triangle{
		NewTriangle(Vertex{0, 0}, Vertex{1, 1}, Vertex{1, 2}),
		NewTriangle(Vertex{1, 1}, Vertex{1, 2}, Vertex{2, 1}),
		NewTriangle(Vertex{1, 2}, Vertex{2, 1}, Vertex{4, 4}),
		NewTriangle(Vertex{2, 1}, Vertex{4, 4}, Vertex{3, 8}),
	}

	st := NewTriangle(
		Vertex{3, 8},
		Vertex{-10, -10},
		Vertex{10, -10},
	)

	result := removeSuperTriangle(arrayOfTrianglesToMap(triangles), st)
	expected := []Triangle{
		NewTriangle(Vertex{0, 0}, Vertex{1, 1}, Vertex{1, 2}),
		NewTriangle(Vertex{1, 1}, Vertex{1, 2}, Vertex{2, 1}),
		NewTriangle(Vertex{1, 2}, Vertex{2, 1}, Vertex{4, 4}),
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
		NewTriangle(Vertex{0, 0}, Vertex{5, 2}, Vertex{2, 5}),
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
		NewTriangle(Vertex{-10, 8}, Vertex{0, 0}, Vertex{5, 17}),
		NewTriangle(Vertex{0, 0}, Vertex{4, 1}, Vertex{5, 17}),
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
		NewTriangle(Vertex{0, 0}, Vertex{0, 8}, Vertex{4, 1}),
		NewTriangle(Vertex{0, 8}, Vertex{4, 1}, Vertex{5, 7}),
		NewTriangle(Vertex{4, 1}, Vertex{8, 2}, Vertex{5, 7}),
		NewTriangle(Vertex{0, 8}, Vertex{5, 7}, Vertex{9, 14}),
		NewTriangle(Vertex{5, 7}, Vertex{8, 2}, Vertex{9, 14}),
	}

	result := triangulate(stars)

	if !sameArrayOFTriangles(result, expected) {
		t.Errorf("Result: %v   expected: %v", result, expected)
	}
}

func BenchmarkTriangulation(b *testing.B) {
	points := generateRandomPoints(1000)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		triangulate(points)
	}
}

// func TestTriangulation(t *testing.T) {
// 	points := generateRandomPoints(1000)
// 	triangulate(points)
// }
