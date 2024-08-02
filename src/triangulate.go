package main

import (
	"fmt"
	"math"
	"sort"
)

type Comparable interface {
	Equals(other interface{}) bool
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Equals(other interface{}) bool {
	if otherVertex, ok := other.(Vertex); ok {
		return v.X == otherVertex.X && v.Y == otherVertex.Y
	}
	return false
}

type Edge struct {
	v0, v1 Vertex
}

func (e Edge) Equals(other interface{}) bool {
	if otherEdge, ok := other.(Edge); ok {
		return e.v0.Equals(otherEdge.v0) && e.v1.Equals(otherEdge.v1) ||
			e.v0.Equals(otherEdge.v1) && e.v1.Equals(otherEdge.v0)
	}
	return false
}

func (e Edge) Hash() string {
	points := []Vertex{e.v0, e.v1}

	sort.Slice(points, func(i, j int) bool {
		if points[i].X == points[j].X {
			return points[i].Y < points[j].Y
		}
		return points[i].X < points[j].X
	})

	return fmt.Sprintf("%v,%v-%v,%v", points[0].X, points[0].Y, points[1].X, points[1].Y)
}

type Triangle struct {
	v0, v1, v2 Vertex
}

func (t Triangle) Equals(other interface{}) bool {
	if t.v0 != other.(Triangle).v0 {
		return false
	}
	if t.v1 != other.(Triangle).v1 {
		return false
	}
	if t.v2 != other.(Triangle).v2 {
		return false
	}
	return true
}

func (t Triangle) inCircumcircle(other Vertex) bool {
	return false
}

func vertexDistance(v1, v2 Vertex) float64 {
	d1 := math.Pow(v1.X-v2.X, 2)
	d2 := math.Pow(v1.Y-v2.Y, 2)

	return math.Sqrt(d1 + d2)
}

func superTriangle(stars []Vertex) Triangle {
	minx, miny, maxx, maxy := math.Inf(1), math.Inf(1), math.Inf(-1), math.Inf(-1)

	for _, star := range stars {
		x := star.X
		y := star.Y

		minx = math.Min(minx, x)
		miny = math.Min(miny, y)
		maxx = math.Max(maxx, x)
		maxy = math.Max(maxy, y)
	}

	dx := maxx - minx
	dy := maxy - miny

	v0 := Vertex{(maxx + minx) / 2, maxy + (dy * 100)}
	v1 := Vertex{minx - (dx * 100), miny - (dy * 100)}
	v2 := Vertex{maxx + (dx * 100), miny - (dy * 100)}

	return Triangle{v0, v1, v2}
}

func circumCenter(tri Triangle) (Vertex, float64) {
	a := tri.v0
	b := tri.v1
	c := tri.v2

	d := 2 * ((a.X * (b.Y - c.Y)) + (b.X * (c.Y - a.Y)) + (c.X * (a.Y - b.Y)))

	centerpoint := Vertex{}

	ax2 := math.Pow(a.X, 2)
	ay2 := math.Pow(a.Y, 2)

	bx2 := math.Pow(b.X, 2)
	by2 := math.Pow(b.Y, 2)

	cx2 := math.Pow(c.X, 2)
	cy2 := math.Pow(c.Y, 2)

	centerpoint.X = ((ax2+ay2)*(b.Y-c.Y) + (bx2+by2)*(c.Y-a.Y) + (cx2+cy2)*(a.Y-b.Y)) / d
	centerpoint.Y = ((ax2+ay2)*(c.X-b.X) + (bx2+by2)*(a.X-c.X) + (cx2+cy2)*(b.X-a.X)) / d

	radius := vertexDistance(centerpoint, a)

	return centerpoint, radius
}

func inCircumcircle(tri Triangle, p Vertex) bool {
	center, radius := circumCenter(tri)
	distance := vertexDistance(p, center)
	return distance <= radius
}

// func uniqueEdge(tri Triangle, badTriangles []Triangle) bool {
// 	edges := [3]Edge{
// 		{tri.v0, tri.v1},
// 		{tri.v1, tri.v2},
// 		{tri.v2, tri.v0},
// 	}

// 	for _, edge := range edges {

// 		for _, other := range badTriangles {
// 			if other == tri {
// 				continue
// 			}

// 			e0 := Edge{other.v0, other.v1}
// 			e1 := Edge{other.v1, other.v2}
// 			e2 := Edge{other.v2, other.v0}

// 			if edge == e0 || edge == e1 || edge == e2 {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

func uniqueEdges(badTriangles []Triangle) []Edge {
	uniques := make(map[string]struct{})
	var uniqueEdges []Edge

	for _, tri := range badTriangles {
		edges := []Edge{
			{tri.v0, tri.v1},
			{tri.v1, tri.v2},
			{tri.v2, tri.v0},
		}

		for _, edge := range edges {
			hash := edge.Hash()
			if _, exists := uniques[hash]; !exists {
				uniques[hash] = struct{}{}
				uniqueEdges = append(uniqueEdges, edge)
			}
		}
	}

	return uniqueEdges
}

// func triangulate(stars []Vertex) []Triangle {

// 	st := superTriangle(stars)

// 	triangulation := make([]Triangle, len(stars))
// 	triangulation = append(triangulation, st)

// 	for _, point := range stars {
// 		badTriangles := make([]Triangle, 1)

// 		for _, tri := range triangulation {
// 			if inCircumcircle(tri, point) {
// 				badTriangles = append(badTriangles, tri)
// 			}
// 		}

// 		polygon := uniqueEdges(badTriangles) // find the boundary of the polygonal hole

// 		newTriangulation := make([]Triangle, len(triangulation))
// 		for _, Tri := range triangulation {
// 			add := true
// 			for _, badTri := range badTriangles {
// 				if badTri == Tri {
// 					add = false
// 					break
// 				}
// 			}

// 			if add {
// 				newTriangulation = append(newTriangulation, Tri)
// 			}
// 		}
// 		triangulation = newTriangulation

// 		for _, e := range polygon {
// 			newTri := Triangle{point, e.v1, e.v2}
// 			triangulation = append(triangulation, newTri)
// 		}

// 	}

// 	// // remove big triangle
// 	// finalTriangulation := make([]Triangle, len(triangulation))
// 	// for _, tri := range triangulation {
// 	// 	edges := [3]Edge{
// 	// 		{tri.v0, tri.v1},
// 	// 		{tri.v1, tri.v2},
// 	// 		{tri.v2, tri.v0},
// 	// 	}

// 	// 	for _, edge := range edges {
// 	// 		if edge == stEdges[0] || edge == stEdges[1] || edge == stEdges[2] {
// 	// 			break
// 	// 		}
// 	// 	}
// 	// }

// 	// return finalTriangulation

// 	return triangulation

// }

// // https://www.gorillasun.de/blog/bowyer-watson-algorithm-for-delaunay-triangulation/
// // https://www.desmos.com/calculator/0waviug7kx
