package controller

import "math"

func distanceFrom(a, b Coordinate) int {
	x := float64(abs(a.x - b.x))
	y := float64(abs(a.y - b.y))
	distance := int(math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2)))
	return distance
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
