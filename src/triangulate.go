package main

import "math"

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

func superTriangle(stars []star) Triangle {

	minx, miny, maxx, maxy := math.Inf(1), math.Inf(1), math.Inf(-1), math.Inf(-1)

	for _, star := range stars {
		x := float64(star.location[0])
		y := float64(star.location[1])

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

func triangulate(stars []star) []Triangle {
	a := make([]Triangle, 1)
	a[0] = Triangle{Vertex{1, 2}, Vertex{1, 2}, Vertex{1, 2}}
	return a
}

// https://www.gorillasun.de/blog/bowyer-watson-algorithm-for-delaunay-triangulation/
