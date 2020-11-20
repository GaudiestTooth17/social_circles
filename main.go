package main

import (
	"github.com/GaudiestTooth17/social_circles_go/controller"
	"github.com/GaudiestTooth17/social_circles_go/model"
)

func main() {
	// emptyConfiguration := controller.NetworkConfig{ReachConfigs: make(map[string]model.ReachConfig),
	// 	NoIsolates: false, GridSize: 0, SaveCSV: false}
	// view.MainMenu(emptyConfiguration)
	configuration := controller.NetworkConfig{
		ReachConfigs: map[string]model.ReachConfig{
			"green":  model.ReachConfig{Quantity: 27710, Reach: 5},
			"blue":   model.ReachConfig{Quantity: 5542, Reach: 10},
			"purple": model.ReachConfig{Quantity: 3695, Reach: 15},
		},
		GridSize:    1223,
		NoIsolates:  false,
		SaveCSV:     false,
		NetworkName: "socnet2-w-coords.txt",
	}
	graph := controller.NewAdjacencyMatrix(configuration)
	controller.SaveAdjacencyList(configuration.NetworkName, graph, true)
}
