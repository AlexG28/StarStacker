package main

import (
	"fmt"
	"image"
)

type star struct {
	id       int
	points   [][2]int
	location [2]int
}

func (s *star) print() string {
	return fmt.Sprintf("ID: %v\n location: (%v, %v)\n\n", s.id, s.location[0], s.location[1])
}

func countStars(img *image.Gray) []star {
	bounds := img.Bounds()
	visited := *createVisited(bounds.Max.X, bounds.Max.Y)
	count := 0
	var stars []star
	for x := range bounds.Max.X {
		for y := range bounds.Max.Y {
			white := img.GrayAt(x, y).Y == 255
			labelled := visited[x][y] != 0

			switch {
			case white && labelled:
				continue
			case white && !labelled:
				count += 1
				points := bfs(x, y, &visited, img, count) // need to do something with this
				starLocation := calculateStarLocation(points)
				newStar := star{id: count, points: points, location: starLocation}
				stars = append(stars, newStar)
				fmt.Printf("res: %v\n", points)
			}
		}
	}
	return stars
}

func bfs(x, y int, visited *[][]int, img *image.Gray, label int) [][2]int {
	var star [][2]int
	star = append(star, [2]int{x, y})

	(*visited)[x][y] = label

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
					(*visited)[neighbour[0]][neighbour[1]] = label
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
	maxX := len(*visited)
	maxY := len((*visited)[0])

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

func createVisited(h, w int) *[][]int {
	visited := make([][]int, h)
	for i := range visited {
		visited[i] = make([]int, w)
	}
	return &visited
}

func calculateStarLocation(points [][2]int) [2]int {
	sumX, sumY := 0, 0
	length := len(points)

	for _, point := range points {
		sumX += point[0]
		sumY += point[1]
	}

	return [2]int{sumX / length, sumY / length}
}
