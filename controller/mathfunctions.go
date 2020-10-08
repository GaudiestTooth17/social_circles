package controller

import (
	"math"

	"github.com/GaudiestTooth17/social_circles_go/model"
)

func distanceFrom(a, b model.Coordinate) int {
	x := float64(abs(a.X - b.X))
	y := float64(abs(a.Y - b.Y))
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
