package main

import (
	"testing"
)

func TestFirst(t *testing.T) {

	triangles1 := []Triangle{
		NewTriangle(Vertex{0, 0}, Vertex{1, 0}, Vertex{1, 1}),
	}

	triangles2 := []Triangle{
		NewTriangle(Vertex{1, 0}, Vertex{2, 0}, Vertex{2, 1}),
	}

	res := findTranslation(triangles1, triangles2)

	if res.horizontal >= 0 && res.vertical >= 0 {
		t.Errorf("it didn't work :(")
	}
}
