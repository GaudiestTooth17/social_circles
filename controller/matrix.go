package controller

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/GaudiestTooth17/social_circles_go/model"
)

type void struct{}

func makeAdjacencyMatrix(grid [][]model.ReachType, noIsolates bool) [][]uint8 {
	// initialize neighbors with the agents that can reach each other
	neighbors := make(map[model.Coordinate]map[model.Coordinate]void)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > 0 {
				agentLoc := model.Coordinate{X: i, Y: j}
				neighbors[agentLoc] = searchForNeighbors(agentLoc, grid)
			}
		}
	}

	// if the noIsolates flag is true remove all the nodes that didn't connect
	// to anything
	if noIsolates {
		for agent, agentsNeighbors := range neighbors {
			if len(agentsNeighbors) == 0 {
				delete(neighbors, agent)
			}
		}
	}

	// create a mapping of agent location to ID
	agentToID := make(map[model.Coordinate]int)
	currentID := 0
	for agent := range neighbors {
		agentToID[agent] = currentID
		currentID++
	}

	numAgents := len(neighbors)
	adjacencyMatrix := make([][]uint8, numAgents)
	for i := 0; i < numAgents; i++ {
		adjacencyMatrix[i] = make([]uint8, numAgents)
	}

	// add an edge between the agents within mutual range
	for agent, aNeighbors := range neighbors {
		for neighbor := range aNeighbors {
			adjacencyMatrix[agentToID[agent]][agentToID[neighbor]] = 1
			adjacencyMatrix[agentToID[neighbor]][agentToID[agent]] = 1
		}
	}

	return adjacencyMatrix
}

// searchForNeighbors creates a set of the coordinates of every agent within mutual range of
// the provided agent
func searchForNeighbors(agentLoc model.Coordinate, grid [][]model.ReachType) map[model.Coordinate]void {
	reach := int(grid[agentLoc.X][agentLoc.Y])
	minX := max(0, agentLoc.X-reach)
	maxX := min(len(grid)-1, agentLoc.X+reach)
	minY := max(0, agentLoc.Y-reach)
	maxY := min(len(grid[0])-1, agentLoc.Y+reach)

	neighbors := make(map[model.Coordinate]void)
	for i := minX; i < maxX; i++ {
		for j := minY; j < maxY; j++ {
			if grid[i][j] > 0 && i != agentLoc.X && j != agentLoc.Y {
				neighborLoc := model.Coordinate{X: i, Y: j}
				distance := distanceFrom(agentLoc, neighborLoc)
				// check to make sure both agents are in range of each other
				if distance <= reach && distance <= int(grid[i][j]) {
					neighbors[neighborLoc] = void{}
				}
			}
		}
	}

	return neighbors
}

// FastSaveMatrix uses "advanced techniques" to speed up write time
func FastSaveMatrix(networkName string, matrix [][]uint8) {
	outFile, err := os.Create(networkName + ".txt")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	var lineBuilder strings.Builder
	lineBuilder.Grow(2 * len(matrix))
	for _, row := range matrix {
		lineBuilder.WriteString(strconv.Itoa(int(row[0])))
		for i := 1; i < len(row); i++ {
			lineBuilder.WriteString(" " + strconv.Itoa(int(row[i])))
		}
		lineBuilder.WriteRune('\n')
		outFile.WriteString(lineBuilder.String())
		lineBuilder.Reset()
	}
}

// SaveAdjacencyList creates a file with the number of nodes and then a list of
// edges separated by newlines
func SaveAdjacencyList(networkName string, matrix [][]uint8) {
	outFile, err := os.Create(networkName + ".txt")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	fmt.Fprintf(outFile, "%d\n", len(matrix))
	for i, row := range matrix {
		for j, value := range row {
			if value > 0 {
				fmt.Fprintf(outFile, "%d %d\n", i, j)
			}
		}
	}
}
