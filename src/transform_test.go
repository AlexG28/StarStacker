package main

import (
	"testing"
)

func TestTranslationBasic(t *testing.T) {

	triangles1 := []Triangle{
		NewTriangle(Vertex{0, 0}, Vertex{1, 0}, Vertex{1, 1}),
	}

	triangles2 := []Triangle{
		NewTriangle(Vertex{1, 0}, Vertex{2, 0}, Vertex{2, 1}),
	}

	res := findTranslation(triangles1, triangles2)

	expected := translation{0.0, 1.0}

	if !res.equalTranslation(expected) {
		t.Errorf("Expected: %v %v    Got: %v %v", expected.horizontal, expected.vertical, res.horizontal, res.vertical)
	}
}

func TestTranslationMedium(t *testing.T) {

	triangles1 := []Triangle{
		NewTriangle(Vertex{0, 0}, Vertex{1, 0}, Vertex{1, 1}),
		NewTriangle(Vertex{1, 0}, Vertex{2, 0}, Vertex{1, 1}),
		NewTriangle(Vertex{1, 1}, Vertex{2, 0}, Vertex{2, 1}),
		NewTriangle(Vertex{0, 0}, Vertex{1, 1}, Vertex{0, 1}),
		NewTriangle(Vertex{0, 1}, Vertex{1, 1}, Vertex{1, 2}),
	}

	triangles2 := []Triangle{
		NewTriangle(Vertex{2, -1}, Vertex{3, -1}, Vertex{3, 0}),
		NewTriangle(Vertex{3, -1}, Vertex{4, -1}, Vertex{3, 0}),
		NewTriangle(Vertex{3, 0}, Vertex{4, -1}, Vertex{4, 0}),
		NewTriangle(Vertex{2, -1}, Vertex{3, 0}, Vertex{2, 0}),
		NewTriangle(Vertex{2, 0}, Vertex{3, 0}, Vertex{3, 1}),
	}

	res := findTranslation(triangles1, triangles2)

	expected := translation{-1.0, 2.0}

	if !res.equalTranslation(expected) {
		t.Errorf("Expected: %v %v    Got: %v %v", expected.horizontal, expected.vertical, res.horizontal, res.vertical)
	}
}
