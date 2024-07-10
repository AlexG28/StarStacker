package main

import (
	"fmt"
	"image"
)

func countStars(img *image.Gray) int {
	bounds := img.Bounds()
	visited := createVisited(bounds.Max.X, bounds.Max.Y)
	count := 0
	for y := range bounds.Max.Y {
		for x := range bounds.Max.X {
			white := img.GrayAt(x, y).Y == 255
			labelled := visited[x][y] != 0

			switch {
			case white && labelled:
				continue
			case white && !labelled:
				res := bfs(x, y, &visited, img)
				fmt.Printf("res: %v\n", res)
			}
		}
	}
	return count
}

func bfs(x, y int, visited *[][]int, img *image.Gray) [][2]int {
	var star [][2]int
	star = append(star, [2]int{x, y})

	(*visited)[y][x] = 1

	queue := make(Queue, 0)
	queue.enqueue([2]int{x, y})

	for !queue.isEmpty() {
		elem, ok := queue.dequeue()
		if ok {
			point := elem.([2]int)
			adjacentPoints := getAdjacentPoints(point[0], point[1], visited, img)

			for _, neighbour := range adjacentPoints {
				visitedNeighbour := (*visited)[neighbour[0]][neighbour[1]] != 0
				if !visitedNeighbour {
					(*visited)[neighbour[0]][neighbour[1]] = 1
					newPoint := [2]int{neighbour[0], neighbour[1]}
					queue.enqueue(newPoint)
					star = append(star, newPoint)
				}
			}
		}
	}

	return star
}

func getAdjacentPoints(x, y int, visited *[][]int, img *image.Gray) [][2]int {
	maxY := len(*visited)
	maxX := len((*visited)[0])

	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	var adjacentPoints [][2]int

	for _, direction := range directions {
		newX, newY := x+direction[0], y+direction[1]

		validX := newX >= 0 && newX < maxX
		validY := newY >= 0 && newY < maxY

		if !validX || !validY {
			continue
		} else {
			if (*visited)[newX][newY] != 0 {
				continue
			}

			if img.GrayAt(newX, newY).Y == 0 {
				continue
			}
		}

		adjacentPoints = append(adjacentPoints, [2]int{newX, newY})
	}

	return adjacentPoints
}

func createVisited(w, h int) [][]int {
	visited := make([][]int, h)
	for i := range visited {
		visited[i] = make([]int, w)
	}
	return visited
}
