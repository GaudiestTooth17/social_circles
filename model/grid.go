package model

import (
	"math/rand"
	"time"
)

type void struct{}

var seed int64 = int64(time.Now().Second())

// Coordinate records a coordinate for an adjacency matrix
type Coordinate struct {
	// X position
	X int
	// Y Position
	Y int
}

// Grid represents a grid of agents
type Grid struct {
	// reaches is a matrix of agent reaches
	reaches [][]ReachType
	// reachToQuantity maps the reach of agents to their location on the grid
	reachToQuantity map[ReachType]uint32
}

// ReachAt returns the reach of the agent at (x, y) or 0 if there is no agent
func (g *Grid) ReachAt(x, y int) ReachType {
	return g.reaches[x][y]
}

// ReachToQuantity returns a copy of the internal map
func (g Grid) ReachToQuantity() map[ReachType]uint32 {
	return g.reachToQuantity
}

// NewGrid uses a Random uniform distribution to randomly place agents on the grid
func NewGrid(agentReaches []ReachType, sideLength int32) Grid {
	r := rand.New(rand.NewSource(seed))
	reaches := make([][]ReachType, sideLength)
	for i := int32(0); i < sideLength; i++ {
		reaches[i] = make([]ReachType, sideLength)
	}

	reachToQuantity := make(map[ReachType]uint32)
	for _, reach := range agentReaches {
		x, y := chooseAvailableSpot(reaches, r)
		reaches[x][y] = reach
		reachToQuantity[reach]++
	}

	return Grid{reaches: reaches, reachToQuantity: reachToQuantity}
}

func chooseAvailableSpot(grid [][]ReachType, r *rand.Rand) (int, int) {
	x := r.Int31n(int32(len(grid)))
	y := r.Int31n(int32(len(grid[0])))
	for grid[x][y] > 0 {
		x = r.Int31n(int32(len(grid)))
		y = r.Int31n(int32(len(grid[0])))
	}
	return int(x), int(y)
}
