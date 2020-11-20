package controller

import (
	"fmt"

	"github.com/GaudiestTooth17/social_circles_go/model"
)

// NetworkConfig stores a configuration for the program to use when making a social network
type NetworkConfig struct {
	ReachConfigs map[string]model.ReachConfig
	GridSize     int32
	NoIsolates   bool
	SaveCSV      bool
	NetworkName  string
}

// NewAdjacencyMatrix creates a new adjacency matrix using the given configuration
func NewAdjacencyMatrix(configuration NetworkConfig) Graph {
	reachConfigs := configuration.makeReachConfigSlice()
	reaches := makeAgentReaches(reachConfigs)
	agentGrid := makeGridRandomUniform(reaches, configuration.GridSize)
	adjacencyMatrix := makeAdjacencyMatrix(agentGrid, configuration.NoIsolates)
	return adjacencyMatrix
}

func (n *NetworkConfig) makeReachConfigSlice() []model.ReachConfig {
	configs := make([]model.ReachConfig, len(n.ReachConfigs))
	i := 0
	for _, config := range n.ReachConfigs {
		configs[i] = config
		i++
	}
	return configs
}

// PrintReaches prints all the reaches in the configuration
func (n *NetworkConfig) PrintReaches() {
	fmt.Printf("%v\n", n.ReachConfigs)
}
