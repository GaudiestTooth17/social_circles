package controller

import (
	"math/rand"
	"time"

	"github.com/GaudiestTooth17/social_circles_go/model"
)

var seed int64 = int64(time.Now().UnixNano())

func makeAgentReaches(configs []model.ReachConfig) []model.ReachType {
	r := rand.New(rand.NewSource(seed))
	totalReaches := uint32(0)
	for _, config := range configs {
		totalReaches += config.Quantity
	}
	reaches := make([]uint8, totalReaches)

	reachIndex := 0
	for _, config := range configs {
		for i := uint32(0); i < config.Quantity; i++ {
			reaches[reachIndex] = config.Reach
			reachIndex++
		}
	}

	r.Shuffle(len(reaches), func(i, j int) {
		reaches[i], reaches[j] = reaches[j], reaches[i]
	})

	return reaches
}

func makeGridRandomUniform(agentReaches []model.ReachType, sideLength int32) [][]model.ReachType {
	r := rand.New(rand.NewSource(seed))
	grid := make([][]model.ReachType, sideLength)
	for i := int32(0); i < sideLength; i++ {
		grid[i] = make([]model.ReachType, sideLength)
	}

	for _, reach := range agentReaches {
		x, y := chooseAvailableSpot(grid, r)
		grid[x][y] = reach
	}

	return grid
}

func chooseAvailableSpot(grid [][]model.ReachType, r *rand.Rand) (int, int) {
	x := r.Int31n(int32(len(grid)))
	y := r.Int31n(int32(len(grid[0])))
	for grid[x][y] > 0 {
		x = r.Int31n(int32(len(grid)))
		y = r.Int31n(int32(len(grid[0])))
	}
	return int(x), int(y)
}
